package Config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "admin"
	DB_PASSWORD = "lele123"
	DB_NAME     = "local1"
	DB_PORT     = 5432

	db  *gorm.DB
	err error
)

func StartPostgresqlDB() *gorm.DB {

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	return db

}
