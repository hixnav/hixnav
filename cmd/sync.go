package cmd

import (
	"gitee.com/wennmu/gopkg.git/doconfig"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"github.com/hixnav/hixnav.git/internal/storage"
)

// SyncPushToCOS manually pushes local data to Tencent COS
func SyncPushToCOS(c *gin.Context) (interface{}, error) {
	driverType := doconfig.GetString("storage_driver")
	if driverType != "sqlite" {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "sync to COS is only supported with sqlite storage driver",
		}
	}

	sync := storage.NewCOSSync(
		GlobalCOS,
		GlobalCOSSecretID,
		GlobalCOSSecretKey,
		doconfig.GetString("cos_sync_bucket"),
		storage.GetDB(),
	)

	result, err := sync.PushToCOS()
	if err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	return result, nil
}

// SyncPullFromCOS manually pulls data from Tencent COS to local SQLite
func SyncPullFromCOS(c *gin.Context) (interface{}, error) {
	driverType := doconfig.GetString("storage_driver")
	if driverType != "sqlite" {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "sync from COS is only supported with sqlite storage driver",
		}
	}

	sync := storage.NewCOSSync(
		GlobalCOS,
		GlobalCOSSecretID,
		GlobalCOSSecretKey,
		doconfig.GetString("cos_sync_bucket"),
		storage.GetDB(),
	)

	result, err := sync.PullFromCOS()
	if err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	return result, nil
}

// SyncExportToFile exports local data to JSON files
func SyncExportToFile(c *gin.Context) (interface{}, error) {
	driverType := doconfig.GetString("storage_driver")
	if driverType != "sqlite" {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "export is only supported with sqlite storage driver",
		}
	}

	exportDir := doconfig.GetString("cos_sync_export_dir")
	if exportDir == "" {
		exportDir = "data/export"
	}

	sync := storage.NewCOSSync(
		GlobalCOS,
		GlobalCOSSecretID,
		GlobalCOSSecretKey,
		doconfig.GetString("cos_sync_bucket"),
		storage.GetDB(),
	)

	if err := sync.ExportToJSONFile(exportDir); err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	return map[string]interface{}{
		"success": true,
		"message": "Data exported to " + exportDir,
	}, nil
}

// SyncImportFromFile imports data from JSON files to local SQLite
func SyncImportFromFile(c *gin.Context) (interface{}, error) {
	driverType := doconfig.GetString("storage_driver")
	if driverType != "sqlite" {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "import is only supported with sqlite storage driver",
		}
	}

	exportDir := doconfig.GetString("cos_sync_export_dir")
	if exportDir == "" {
		exportDir = "data/export"
	}

	sync := storage.NewCOSSync(
		GlobalCOS,
		GlobalCOSSecretID,
		GlobalCOSSecretKey,
		doconfig.GetString("cos_sync_bucket"),
		storage.GetDB(),
	)

	result, err := sync.ImportFromJSONFile(exportDir)
	if err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	return result, nil
}

// GetSyncStatus returns the current sync configuration status
func GetSyncStatus(c *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"storage_driver":    doconfig.GetString("storage_driver"),
		"sqlite_path":       doconfig.GetString("sqlite_path"),
		"cos_sync_enabled":  doconfig.GetString("cos_sync_enabled"),
		"cos_sync_interval": doconfig.GetString("cos_sync_interval"),
		"cos_sync_bucket":   doconfig.GetString("cos_sync_bucket"),
	}, nil
}
