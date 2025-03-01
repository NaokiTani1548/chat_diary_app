package handlers

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getDBConnection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	dsn := os.Getenv("DB_ROLE") + ":" + os.Getenv("DB_PASS") + "@/" + os.Getenv("DB_NAME")
	return sql.Open("mysql", dsn)
}

func InitDB() {
	DB_Connection, err := getDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer DB_Connection.Close()

	if err = DB_Connection.Ping(); err != nil {
		log.Fatalf("DB ping error: %v", err)
	}

	cmd := `INSERT INTO author (author_id) VALUES (1);`
	DB_Connection.Exec(cmd)
}
