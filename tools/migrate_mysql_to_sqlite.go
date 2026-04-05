package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitee.com/wennmu/gopkg.git/doconfig"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 模型定义
type User struct {
	Id       int64
	Account  string
	Passwd   string
	CreateAt int64
}

func (User) TableName() string { return "users" }

type Nav struct {
	Id       int64
	Cate     int64
	Catename string
	Name     string
	Logo     string
	Desc     string
	Url      string
	Uid      int64
}

func (Nav) TableName() string { return "navs" }

type Link struct {
	Id       int64
	Type     int64
	Catename string
	Name     string
	Logo     string
	Url      string
	Uid      int64
}

func (Link) TableName() string { return "links" }

type Account struct {
	ID       int64
	Sitename string
	Siteurl  string
	Name     string
	Password string
	Createat int64
	Uid      int64
}

func (Account) TableName() string { return "accounts" }

func main() {
	// 初始化配置
	if err := doconfig.InitConfigs(); err != nil {
		log.Fatal("加载配置失败:", err)
	}

	mysqlDSN := doconfig.GetString("DB")
	sqlitePath := doconfig.GetString("sqlite_path")
	if sqlitePath == "" {
		sqlitePath = "data/hixnav.db"
	}

	if mysqlDSN == "" {
		log.Fatal("MySQL DSN 为空，请检查 config.yaml 中的 DB 配置")
	}

	// 连接 MySQL
	mysqlDB, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("连接 MySQL 失败:", err)
	}
	fmt.Println("✓ MySQL 连接成功")

	// 确保 SQLite 目录存在
	dir := filepath.Dir(sqlitePath)
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}

	// 连接 SQLite
	sqliteDB, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		log.Fatal("连接 SQLite 失败:", err)
	}
	fmt.Println("✓ SQLite 连接成功:", sqlitePath)

	// 自动建表
	sqliteDB.AutoMigrate(&User{}, &Nav{}, &Link{}, &Account{})
	fmt.Println("✓ SQLite 表结构创建完成")

	// 迁移数据
	tables := []struct {
		name string
		data interface{}
	}{
		{name: "users", data: &[]User{}},
		{name: "navs", data: &[]Nav{}},
		{name: "links", data: &[]Link{}},
		{name: "accounts", data: &[]Account{}},
	}

	totalRecords := 0
	for _, t := range tables {
		// 从 MySQL 读取
		result := mysqlDB.Table(t.name).Find(t.data)
		if result.Error != nil {
			fmt.Printf("  ⚠ 读取 MySQL 表 %s 失败: %v\n", t.name, result.Error)
			continue
		}

		var count int
		switch v := t.data.(type) {
		case *[]User:
			count = len(*v)
		case *[]Nav:
			count = len(*v)
		case *[]Link:
			count = len(*v)
		case *[]Account:
			count = len(*v)
		}

		if count == 0 {
			fmt.Printf("  - 表 %s: 无数据，跳过\n", t.name)
			continue
		}

		// 写入 SQLite（先清空再插入，确保数据一致）
		sqliteDB.Exec(fmt.Sprintf("DELETE FROM %s", t.name))
		createResult := sqliteDB.Table(t.name).Create(t.data)
		if createResult.Error != nil {
			fmt.Printf("  ✗ 写入 SQLite 表 %s 失败: %v\n", t.name, createResult.Error)
			continue
		}

		totalRecords += count
		fmt.Printf("  ✓ 表 %s: 迁移 %d 条记录\n", t.name, count)
	}

	fmt.Printf("\n迁移完成！共迁移 %d 条记录到 %s\n", totalRecords, sqlitePath)
}
