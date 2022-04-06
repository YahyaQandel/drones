package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectToDatabase(args ...string) (*gorm.DB, error) {
	dbConfig := getDBConfigurations()
	if len(args) > 0 {
		dbConfig = args[0]
	}
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "dronetask.",
		},
	})
	return db, err
}

func getDBConfigurations() string {
	dbConfig := os.Getenv("DB_CONNECTION_STRING")
	log.Println(dbConfig)
	return dbConfig
}
