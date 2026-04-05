package cmd

import (
	"encoding/base64"
	"fmt"
	"gitee.com/wennmu/gopkg.git/doconfig"
	"gitee.com/wennmu/gopkg.git/dologger"
	"github.com/gin-gonic/gin"
	"github.com/hixnav/hixnav.git/internal/e"
	"github.com/hixnav/hixnav.git/internal/errcode"
	"github.com/hixnav/hixnav.git/internal/storage"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func InitDB() error {
	driverType := doconfig.GetString("storage_driver")
	if driverType == "" {
		driverType = "mysql"
	}

	var err error
	switch storage.DriverType(driverType) {
	case storage.DriverSQLite:
		sqlitePath := doconfig.GetString("sqlite_path")
		if sqlitePath == "" {
			sqlitePath = "data/hixnav.db"
		}
		err = storage.Register(storage.NewSQLiteDriver(sqlitePath))
		// Auto-migrate tables for SQLite
		if err == nil {
			db := storage.GetDB()
			if db != nil {
				err = db.AutoMigrate(&User{}, &Nav{}, &Link{}, &Account{})
				if err != nil {
					dologger.Error("sqlite auto-migrate failed:", err.Error())
					return err
				}
			}
		}
	case storage.DriverMySQL:
		fallthrough
	default:
		dsn := doconfig.GetString("DB")
		if dsn == "" {
			dologger.Warn("DB DSN is empty, falling back to sqlite")
			sqlitePath := doconfig.GetString("sqlite_path")
			if sqlitePath == "" {
				sqlitePath = "data/hixnav.db"
			}
			err = storage.Register(storage.NewSQLiteDriver(sqlitePath))
			// Auto-migrate tables for SQLite fallback
			if err == nil {
				db := storage.GetDB()
				if db != nil {
					err = db.AutoMigrate(&User{}, &Nav{}, &Link{}, &Account{})
					if err != nil {
						dologger.Error("sqlite auto-migrate failed:", err.Error())
						return err
					}
				}
			}
		} else {
			err = storage.Register(storage.NewMySQLDriver(dsn))
		}
	}

	if err != nil {
		dologger.Error("init db failed:", err.Error())
		return err
	}

	// Set the global DB reference for backward compatibility
	GlobalDB = storage.GetDB()

	return nil
}

//数据库迁移
type DbConfig struct {
	Addr      string
	Dbname    string
	Account   string
	Password  string
	IsMigrate bool
}

func MigrateDB(c *gin.Context) (interface{}, error) {
	var dbconf DbConfig
	if err := c.ShouldBindJSON(&dbconf); err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "invalid params",
		}
	}
	doconfig.Set("DB", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbconf.Account, dbconf.Password, dbconf.Addr, dbconf.Dbname))

	if err := createDir("./config"); err != nil {
		return "", e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}
	doconfig.WriteConfigAs("config/config.yaml")

	//初始化数据库
	if err := InitDB(); err != nil {
		return "", e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}

	if !dbconf.IsMigrate {
		return "", nil
	}
	//创建数据库表
	db := storage.GetDB()
	if db == nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "db not initialized",
		}
	}
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Nav{}, &Link{}, &Account{})
	if err != nil {
		log.Println(err.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  "init failed",
		}
	}
	//内置用户
	res := db.Create(&User{
		Account:  defaultUsername,
		Passwd:   base64.StdEncoding.EncodeToString([]byte(defaultUserPasswd)),
		CreateAt: time.Now().Unix(),
	})
	if res.Error != nil {
		log.Println(res.Error.Error())
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  res.Error.Error(),
		}
	}
	return "", nil
}

//检查数据库可用性
func DbTest(c *gin.Context) (interface{}, error) {
	db := storage.GetDB()
	if db == nil {
		return "", e.AppError{Code: -1, Msg: "db not initialized"}
	}
	// Try a simple query to test connectivity
	var result int
	if err := db.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return "", e.AppError{Code: -1, Msg: "db invalid. Please modify the database configuration information."}
	}
	return "", nil
}

type Table struct {
	TableName string `gorm:"column:TABLE_NAME"`
}

func ExportDatabase(c *gin.Context) (interface{}, error) {
	// 生成SQL文件名
	filename := "hixnav.sql"

	// 准备写入SQL文件
	file, err := os.Create(filename)
	if err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}
	defer file.Close()

	// Get the active driver type
	driverType := doconfig.GetString("storage_driver")
	if driverType == "" {
		driverType = "mysql"
	}

	if storage.DriverType(driverType) == storage.DriverMySQL {
		// MySQL-specific export using information_schema
		var tables []Table
		if err := storage.GetDB().Raw("SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?", "mydb").Scan(&tables).Error; err != nil {
			return nil, e.AppError{
				Code: errcode.ERROR,
				Msg:  err.Error(),
			}
		}

		// 遍历表名并导出数据
		for _, table := range tables {
			tableName := table.TableName

			// 生成导出数据的SQL语句
			stmt := fmt.Sprintf("SELECT * INTO OUTFILE '%s' FROM %s", tableName+".sql", tableName)

			// 执行导出SQL语句
			if err := storage.GetDB().Exec(stmt).Error; err != nil {
				return nil, e.AppError{
					Code: errcode.ERROR,
					Msg:  err.Error(),
				}
			}

			// 读取导出的SQL文件内容
			sqlBytes, err := ioutil.ReadFile(tableName + ".sql")
			if err != nil {
				return nil, e.AppError{
					Code: errcode.ERROR,
					Msg:  err.Error(),
				}
			}

			// 将SQL文件内容写入总的SQL文件
			_, err = file.Write(sqlBytes)
			if err != nil {
				return nil, e.AppError{
					Code: errcode.ERROR,
					Msg:  err.Error(),
				}
			}

			// 删除单独的SQL文件
			if err := os.Remove(tableName + ".sql"); err != nil {
				return nil, e.AppError{
					Code: errcode.ERROR,
					Msg:  err.Error(),
				}
			}
		}
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, e.AppError{
			Code: errcode.ERROR,
			Msg:  err.Error(),
		}
	}
	return map[string]interface{}{
		"file":          data,
		"contentType":   "application/octet-stream",
		"contentLength": len(data),
	}, nil
}
