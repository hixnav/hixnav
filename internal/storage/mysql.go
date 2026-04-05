package storage

import (
	"gitee.com/wennmu/gopkg.git/dologger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDriver implements StorageDriver for MySQL
type MySQLDriver struct {
	dsn string
	db  *gorm.DB
}

// NewMySQLDriver creates a new MySQL driver instance
func NewMySQLDriver(dsn string) *MySQLDriver {
	return &MySQLDriver{dsn: dsn}
}

func (d *MySQLDriver) Name() DriverType {
	return DriverMySQL
}

func (d *MySQLDriver) Init() error {
	var err error
	d.db, err = gorm.Open(mysql.Open(d.dsn), &gorm.Config{})
	if err != nil {
		dologger.Error("mysql connect failed:", err.Error())
		return err
	}
	dologger.Info("mysql connected successfully")
	return nil
}

func (d *MySQLDriver) DB() *gorm.DB {
	return d.db
}

func (d *MySQLDriver) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
