package db_util

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Name string
	User string
	Pass string
	Host string
}

func InitDb(cfg MysqlConfig, models []interface{}) *gorm.DB {
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Name,
	)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(models...)
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	return db
}

func InitSqlite(dbFile string, models []interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		err = fmt.Errorf("sqlite.Open(%v): %w", dbFile, err)
		return db
	}

	db.AutoMigrate(models...)

	return db
}
