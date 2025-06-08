package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Connected to database using pgx!")
	return db
}
