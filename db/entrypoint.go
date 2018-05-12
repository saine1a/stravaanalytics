package db

import (
	"database/sql"
	"fmt"
	"os"
	// Bring in the MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DBaccess - DB access object
type DBaccess struct {
}

// Init - create and construct a DBaccess object
func Init() *DBaccess {

	// Connect to DB
	dbInfo := os.Getenv("STRAVA_DB_STRING")

	if dbInfo == "" {
		fmt.Println("Need STRAVA_DB_STRING to contain the MySql access string")
		os.Exit(-1)

	}

	db, err := sql.Open("mysql", dbInfo)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	// Check tables exist

	return &DBaccess{}
}
