package database

import (
	"fmt"

    "github.com/bismarkanes/web-go/infrastructure/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type ping struct {
	ID   int
	Name string
}

// New main database config
func NewDBConnection() *gorm.DB {
    Config := config.NewConfig()

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        Config.DBHost,
        Config.DBUsername,
        Config.DBPassword,
        Config.DBName,
        Config.DBPort,
    )

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
