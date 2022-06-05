package Common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitMysqlConnection() {
	dsn := "root:107058seed00UC@tcp(127.0.0.1:3306)/simple_tik_tok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second)
}

func MysqlConnection() *gorm.DB {
	dsn := "root:107058seed00UC@tcp(127.0.0.1:3306)/simple_tik_tok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
