package database

import (
	"Restapi/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is database connection
var DB *gorm.DB = Dbcon()

// Dbcon is a function for connecting to a database.
func Dbcon() *gorm.DB {
	// config.Envload()

	//Production
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database!")
	db.AutoMigrate(&models.User{}, &models.BarberProfile{}, &models.SalonService{})
	db.AutoMigrate(&models.ServiceList{}, &models.ServiceMetaData{}, &models.UserNotification{})
	// db.Migrator().DropTable(&models.User{}, &models.BarberProfile{}, &models.SalonService{})
	fmt.Println("Database Migration Completed!")
	return db
}
