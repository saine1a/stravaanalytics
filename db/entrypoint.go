package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	// Bring in the MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/saine1a/stravaanalytics/stravaaccess"
)

// DBaccess - DB access object
type DBaccess struct {
	db *sql.DB
}

// Init - create and construct a DBaccess object
func Init() *DBaccess {

	// Connect to DB
	dbInfo := os.Getenv("STRAVA_DB_STRING")

	if dbInfo == "" {
		fmt.Println("Need STRAVA_DB_STRING to contain the MySql access string")
		os.Exit(-1)

	}

	theDb, err := sql.Open("mysql", dbInfo)

	if err != nil {
		panic(err.Error())
	}

	err = theDb.Ping()

	if err != nil {
		panic(err.Error())
	}

	return &DBaccess{db: theDb}
}

func sqlName(name string) string {
	newName := strings.Replace(name, " ", "_", -1)
	newName = strings.Replace(newName, ".", "_", -1)
	return newName
}

func (obj *DBaccess) execSQL(sql string) sql.Result {
	result, err := obj.db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	return result

}

// StoreActivities - store a set of activities in the database
func (obj *DBaccess) StoreActivities(club stravaaccess.Club, activities *([]stravaaccess.SummaryActivity)) {

	// Create schema

	obj.createSchema(club, activities)
}
