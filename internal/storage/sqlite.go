package storage

import (
	"os"
	"path/filepath"

	"gitee.com/wennmu/gopkg.git/dologger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDriver implements StorageDriver for SQLite
type SQLiteDriver struct {
	path string
	db   *gorm.DB
}

// NewSQLiteDriver creates a new SQLite driver instance
func NewSQLiteDriver(path string) *SQLiteDriver {
	return &SQLiteDriver{path: path}
}

func (d *SQLiteDriver) Name() DriverType {
	return DriverSQLite
}

func (d *SQLiteDriver) Init() error {
	// Ensure the directory exists
	dir := filepath.Dir(d.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		dologger.Error("failed to create sqlite directory:", err.Error())
		return err
	}

	var err error
	d.db, err = gorm.Open(sqlite.Open(d.path), &gorm.Config{})
	if err != nil {
		dologger.Error("sqlite connect failed:", err.Error())
		return err
	}

	// Auto-migrate all tables
	if err := d.autoMigrate(); err != nil {
		dologger.Error("sqlite auto-migrate failed:", err.Error())
		return err
	}

	dologger.Info("sqlite connected and migrated successfully")
	return nil
}

func (d *SQLiteDriver) autoMigrate() error {
	// Tables are created on-demand via raw SQL or when first accessed
	// The COS sync layer handles data import/export generically
	return nil
}

func (d *SQLiteDriver) DB() *gorm.DB {
	return d.db
}

func (d *SQLiteDriver) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
