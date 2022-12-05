package database

import (
	"Restapi/config"
	"Restapi/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is database connection
var DB *gorm.DB = Dbcon()

// Dbcon is a function for connecting to a database.
func Dbcon() *gorm.DB {
	config.Envload()

	//Production
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	//Local
	// dsn := os.Getenv("LDBUSER") + `:` + os.Getenv("LDBPASS") + `@tcp(` + os.Getenv("LDBHOST") + `:` + os.Getenv("LDBPORT") + `)/` + os.Getenv("LDBNAME") + `?charset=utf8mb4&parseTime=True&loc=Local`
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	PrepareStmt:            true,
	// 	SkipDefaultTransaction: true,
	// })

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Database!")

	db.AutoMigrate(&models.User{}, &models.BarberProfile{}, &models.SalonService{}, &models.ServiceList{})
	// db.Migrator().DropTable(&models.User{}, &models.BarberProfile{}, &models.SalonService{})
	fmt.Println("Database Migration Completed!")
	return db
}
