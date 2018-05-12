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

// StoreActivities - store a set of activities in the database
func (obj *DBaccess) StoreActivities(club stravaaccess.Club, activities *([]stravaaccess.SummaryActivity)) {

	// First drop the table, in case it already exists

	tableName := fmt.Sprintf("Activity_%s_%d", club.Name, club.ID)

	dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", sqlName(tableName))

	_, err := obj.db.Exec(dropStmt)

	if err != nil {
		panic(err.Error())
	}

	// Now create the table

	createTableStmt := fmt.Sprintf("CREATE TABLE `%s` (`activityName` VARCHAR(256), `rider` VARCHAR(256), `distance` INT(10))", sqlName(tableName))

	fmt.Println(createTableStmt)

	_, err = obj.db.Exec(createTableStmt)

	if err != nil {
		panic(err.Error())
	}

	// New insert the values

	insertStmt, err := obj.db.Prepare(fmt.Sprintf("INSERT %s SET activityName=?,rider=?,distance=?", sqlName(tableName)))

	if err != nil {
		panic(err.Error())
	}

	for a := range *activities {

		activity := (*activities)[a]

		insertStmt.Exec(activity.Name, fmt.Sprintf("%s %s", activity.Athlete.FirstName, activity.Athlete.LastName), activity.Distance)

		if err != nil {
			panic(err.Error())
		}
	}
}
