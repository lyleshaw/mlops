package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DSN")
	if os.Getenv("IS_DEV") == "true" {
		dsn = "test:test@tcp(localhost:3306)/test?parseTime=true"
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}
