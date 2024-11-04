package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/resource/internal/infrastructure/config"
	"github.com/resource/internal/infrastructure/http"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf := config.NewConfig()

	dsn := conf.DatabaseDSN

	if dsn == "" {
		log.Fatalf("DATABASE_DSN environment variable is not set")
	}
	log.Println("value of dsn", dsn)
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Connected to the database")
	router := http.NewRouter()

	log.Println("Starting server on :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Could not listen on :%v: %v\n", conf.ServerPort, err)
	}

}
