package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println("Connecting...")

	// TODO: use env vars
	dsn := "host=localhost user=postgres password=password dbname=auditDB port=5432 sslmode=disable"
	var db, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(fmt.Sprintf("Error init DB: %s\n", error.Error()))
	} else {
		fmt.Println("Success DB")
	}

	DB = db
	// Auto migration
	DB.AutoMigrate(&Audit{})
	println(DB)
}
