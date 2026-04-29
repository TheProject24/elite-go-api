package main

import (
	"log"
	"net/http"
	"database/sql"

	"github.com/jmoiron/sqlx"
	- "github.com/lib/pq"

	"github.com/TheProject24/elite-api/config"
	"github.com/TheProject24/elite-api/handlers"
	"github.com/TheProject24/elite-api/middlewares"
)

import (
	
)

var DB *sqlx.DB


func initDB(dbPassword string) {
	dsn := "user=user password=" + dbPassword + "dbname=trading_bot sslmode = disable"

	var err errorDB, err =sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("CRITICAL ERROR: Failed to connect to the database: ". err)
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)

	log.Println("Successfully connected to the PostgreSQL database")
}


func main() {

	cfg := LoadConfig()
	log.Printf("Booting up Elite API with secure config on port %s...", cfg.Port)

	initDB(cfg.DBPassword)

	defer DB.Close()

	http.HandleFunc("/register", middlewares.AuthMiddleware(handlers.RegisterUser))

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}