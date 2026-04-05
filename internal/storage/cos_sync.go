package storage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitee.com/wennmu/gopkg.git/dologger"
	"github.com/tencentyun/cos-go-sdk-v5"
	"gorm.io/gorm"
)

// COSSync handles synchronization between local SQLite and Tencent COS
type COSSync struct {
	cosURL     string
	secretID   string
	secretKey  string
	bucket     string
	localDB    *gorm.DB
	cosClient  *cos.Client
}

// NewCOSSync creates a new COS sync instance
func NewCOSSync(cosURL, secretID, secretKey, bucket string, localDB *gorm.DB) *COSSync {
	return &COSSync{
		cosURL:    cosURL,
		secretID:  secretID,
		secretKey: secretKey,
		bucket:    bucket,
		localDB:   localDB,
	}
}

// InitCOSClient initializes the COS client
func (s *COSSync) InitCOSClient() error {
	var cosURL string
	if strings.Contains(s.cosURL, "%s") {
		cosURL = fmt.Sprintf(s.cosURL, s.bucket)
	} else {
		cosURL = s.cosURL
	}
	// 确保 URL 以 / 结尾
	if !strings.HasSuffix(cosURL, "/") {
		cosURL += "/"
	}

	u, err := url.Parse(cosURL)
	if err != nil {
		return fmt.Errorf("parse cos url failed: %w", err)
	}
	b := &cos.BaseURL{BucketURL: u}
	s.cosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  s.secretID,
			SecretKey: s.secretKey,
		},
	})
	return nil
}

// TableExport represents a table's exported data
type TableExport struct {
	TableName  string          `json:"table_name"`
	Records    json.RawMessage `json:"records"`
	ExportedAt string          `json:"exported_at"`
}

// SyncResult holds the result of a sync operation
type SyncResult struct {
	Success       bool     `json:"success"`
	Message       string   `json:"message"`
	RecordsSynced int      `json:"records_synced"`
	TablesSynced  []string `json:"tables_synced"`
}

// tableDef defines a table for export/import operations
type tableDef struct {
	name   string
	idCol  string // primary key column name
}

// tablesToSync lists all tables that should be synced
var tablesToSync = []tableDef{
	{name: "users", idCol: "id"},
	{name: "navs", idCol: "id"},
	{name: "links", idCol: "id"},
	{name: "accounts", idCol: "id"},
}

// PushToCOS exports all tables from local SQLite and uploads to COS
func (s *COSSync) PushToCOS() (*SyncResult, error) {
	result := &SyncResult{
		TablesSynced: []string{},
	}

	if s.cosClient == nil {
		if err := s.InitCOSClient(); err != nil {
			return result, fmt.Errorf("init cos client failed: %w", err)
		}
	}

	for _, table := range tablesToSync {
		// Query all records as raw JSON
		var records []map[string]interface{}
		if err := s.localDB.Table(table.name).Find(&records).Error; err != nil {
			dologger.Error(fmt.Sprintf("export table %s failed: %v", table.name, err))
			continue
		}

		// Create export payload
		export := TableExport{
			TableName:  table.name,
			Records:    mustMarshal(records),
			ExportedAt: time.Now().Format(time.RFC3339),
		}

		jsonData, err := json.Marshal(export)
		if err != nil {
			dologger.Error(fmt.Sprintf("marshal table %s failed: %v", table.name, err))
			continue
		}

		// Upload to COS
		objectKey := fmt.Sprintf("db_export/%s.json", table.name)
		reader := bytes.NewReader(jsonData)
		_, err = s.cosClient.Object.Put(context.Background(), objectKey, reader, nil)
		if err != nil {
			dologger.Error(fmt.Sprintf("upload table %s to COS failed: %v", table.name, err))
			continue
		}

		result.RecordsSynced += len(records)
		result.TablesSynced = append(result.TablesSynced, table.name)
		dologger.Info(fmt.Sprintf("pushed table %s to COS: %d records", table.name, len(records)))
	}

	result.Success = len(result.TablesSynced) > 0
	if result.Success {
		result.Message = fmt.Sprintf("Successfully synced %d tables (%d records) to COS", len(result.TablesSynced), result.RecordsSynced)
	} else {
		result.Message = "No tables were synced to COS"
	}

	return result, nil
}

// PullFromCOS downloads data from COS and merges into local SQLite (upsert by ID, never overwrite existing)
func (s *COSSync) PullFromCOS() (*SyncResult, error) {
	result := &SyncResult{
		TablesSynced: []string{},
	}

	if s.cosClient == nil {
		if err := s.InitCOSClient(); err != nil {
			return result, fmt.Errorf("init cos client failed: %w", err)
		}
	}

	for _, table := range tablesToSync {
		objectKey := fmt.Sprintf("db_export/%s.json", table.name)

		// Download from COS
		resp, err := s.cosClient.Object.Get(context.Background(), objectKey, nil)
		if err != nil {
			dologger.Error(fmt.Sprintf("download table %s from COS failed: %v", table.name, err))
			continue
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			dologger.Error(fmt.Sprintf("read table %s response failed: %v", table.name, err))
			continue
		}

		// Parse the export
		var export TableExport
		if err := json.Unmarshal(bodyBytes, &export); err != nil {
			dologger.Error(fmt.Sprintf("unmarshal table %s failed: %v", table.name, err))
			continue
		}

		// Parse records
		var records []map[string]interface{}
		if err := json.Unmarshal(export.Records, &records); err != nil {
			dologger.Error(fmt.Sprintf("unmarshal table %s records failed: %v", table.name, err))
			continue
		}

		// Upsert records by ID
		count := 0
		for _, record := range records {
			if err := s.upsertRecord(table.name, table.idCol, record); err != nil {
				dologger.Error(fmt.Sprintf("upsert %s record failed: %v", table.name, err))
				continue
			}
			count++
		}

		result.RecordsSynced += count
		result.TablesSynced = append(result.TablesSynced, table.name)
		dologger.Info(fmt.Sprintf("pulled table %s from COS: %d records", table.name, count))
	}

	result.Success = len(result.TablesSynced) > 0
	if result.Success {
		result.Message = fmt.Sprintf("Successfully synced %d tables (%d records) from COS", len(result.TablesSynced), result.RecordsSynced)
	} else {
		result.Message = "No tables were synced from COS"
	}

	return result, nil
}

// upsertRecord inserts or updates a record by ID (generic, works with any table)
func (s *COSSync) upsertRecord(tableName, idCol string, record map[string]interface{}) error {
	idVal, ok := record[idCol]
	if !ok {
		return fmt.Errorf("missing %s column", idCol)
	}

	// Check if record exists
	var count int64
	if err := s.localDB.Table(tableName).Where(fmt.Sprintf("%s = ?", idCol), idVal).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		// Record doesn't exist, insert it
		return s.localDB.Table(tableName).Create(record).Error
	}

	// Record exists, update it
	return s.localDB.Table(tableName).Where(fmt.Sprintf("%s = ?", idCol), idVal).Updates(record).Error
}

// ExportToJSONFile exports all tables to local JSON files (for backup/manual upload)
func (s *COSSync) ExportToJSONFile(exportDir string) error {
	if err := os.MkdirAll(exportDir, 0755); err != nil {
		return fmt.Errorf("create export directory failed: %w", err)
	}

	for _, table := range tablesToSync {
		var records []map[string]interface{}
		if err := s.localDB.Table(table.name).Find(&records).Error; err != nil {
			dologger.Error(fmt.Sprintf("export table %s failed: %v", table.name, err))
			continue
		}

		export := TableExport{
			TableName:  table.name,
			Records:    mustMarshal(records),
			ExportedAt: time.Now().Format(time.RFC3339),
		}

		jsonData, err := json.MarshalIndent(export, "", "  ")
		if err != nil {
			dologger.Error(fmt.Sprintf("marshal table %s failed: %v", table.name, err))
			continue
		}

		filePath := filepath.Join(exportDir, fmt.Sprintf("%s.json", table.name))
		if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
			dologger.Error(fmt.Sprintf("write table %s file failed: %v", table.name, err))
			continue
		}

		dologger.Info(fmt.Sprintf("exported table %s to %s", table.name, filePath))
	}

	return nil
}

// ImportFromJSONFile imports data from local JSON files into SQLite (upsert by ID)
func (s *COSSync) ImportFromJSONFile(exportDir string) (*SyncResult, error) {
	result := &SyncResult{
		TablesSynced: []string{},
	}

	for _, table := range tablesToSync {
		filePath := filepath.Join(exportDir, fmt.Sprintf("%s.json", table.name))
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			dologger.Error(fmt.Sprintf("read file %s failed: %v", filePath, err))
			continue
		}

		var export TableExport
		if err := json.Unmarshal(fileData, &export); err != nil {
			dologger.Error(fmt.Sprintf("unmarshal file %s failed: %v", filePath, err))
			continue
		}

		var records []map[string]interface{}
		if err := json.Unmarshal(export.Records, &records); err != nil {
			dologger.Error(fmt.Sprintf("unmarshal table %s records failed: %v", table.name, err))
			continue
		}

		count := 0
		for _, record := range records {
			if err := s.upsertRecord(table.name, table.idCol, record); err != nil {
				continue
			}
			count++
		}

		result.RecordsSynced += count
		result.TablesSynced = append(result.TablesSynced, table.name)
	}

	result.Success = len(result.TablesSynced) > 0
	return result, nil
}

func mustMarshal(v interface{}) json.RawMessage {
	data, _ := json.Marshal(v)
	return data
}
