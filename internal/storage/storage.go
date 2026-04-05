package storage

import (
	"gitee.com/wennmu/gopkg.git/dologger"
	"gorm.io/gorm"
)

// DriverType represents the type of storage driver
type DriverType string

const (
	DriverMySQL  DriverType = "mysql"
	DriverSQLite DriverType = "sqlite"
)

// StorageDriver defines the interface for database storage backends
type StorageDriver interface {
	// Name returns the driver name
	Name() DriverType
	// Init initializes the database connection
	Init() error
	// DB returns the underlying *gorm.DB instance
	DB() *gorm.DB
	// Close closes the database connection
	Close() error
}

// GlobalStorage holds the active storage driver instance
var GlobalStorage StorageDriver

// Register sets the global storage driver
func Register(driver StorageDriver) error {
	if err := driver.Init(); err != nil {
		dologger.Error("storage driver init failed:", err.Error())
		return err
	}
	GlobalStorage = driver
	dologger.Info("storage driver registered:", string(driver.Name()))
	return nil
}

// GetDB returns the global *gorm.DB from the active storage driver
func GetDB() *gorm.DB {
	if GlobalStorage == nil {
		dologger.Error("storage driver not initialized")
		return nil
	}
	return GlobalStorage.DB()
}

// SyncDirection represents the direction of COS sync
type SyncDirection string

const (
	SyncPush SyncDirection = "push" // local -> COS
	SyncPull SyncDirection = "pull" // COS -> local
)
