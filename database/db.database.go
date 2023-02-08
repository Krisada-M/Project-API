package database

import (
	"Restapi/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is database connection
var DB *gorm.DB = Dbcon()

// Dbcon is a function for connecting to a database.
func Dbcon() *gorm.DB {
	if gin.Mode() != "release" {
		config.Envload()
	}

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
	if gin.Mode() != "release" {
		fmt.Println("Database Migration Completed!")
	}
	// db.Migrator().DropTable(&models.User{}, &models.BarberProfile{}, &models.SalonService{})
	return db
}
