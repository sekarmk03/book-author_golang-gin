package database

import (
	"book-author/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDb   = "bookauthor"
	PostgresUser = "postgres"
	PostgresPass = "skrmk372"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPass, PostgresDb,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err := db.AutoMigrate(entity.User{}, entity.Author{}, entity.Book{}, entity.Publisher{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
