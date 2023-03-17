package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/samuelsih/mcsvc/authentication/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const port = "80"

var maxDBRetry = 10

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Connection is nil")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	log.Printf("Starting authentication service on %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres DB not ready...", err)
			maxDBRetry--

			if maxDBRetry == 0 {
				log.Println(err)
				return nil
			}
		} else {
			log.Println("Connect to Postgres")
			return conn
		}

		log.Println("Backing of for 2 sec.")
		time.Sleep(2 * time.Second)
		continue
	}

}
