package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Db_Connect() {
	var err error
	dsn := "host=ep-tiny-snowflake-44902790.ap-southeast-1.postgres.vercel-storage.com user=default password=IuyR5q4sPAZc dbname=verceldb port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		panic("Faild to connect DB. 💪🧵🧵")
	}
}
