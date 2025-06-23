package database

import (
	"fmt"
	"hello_bets/pkg/model"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildConnectionString(user, password, host, port, dbName string) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbName, port,
	)
}

func Connect(user, password, host, port, dbName string) (*gorm.DB, error) {

	dsn := buildConnectionString(user, password, host, port, dbName)
	attempts := 5
	connection := &gorm.DB{}
	var err error
	for {
		connection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			if attempts == 0 {
				log.Println("Error retry database: ", err)
				return nil, err
			} else {
				attempts--
				time.Sleep(2 * time.Second)
			}
		} else {
			break
		}
	}
	return connection, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	return err
}
