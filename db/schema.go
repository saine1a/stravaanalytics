package db

import (
	"fmt"
	"reflect"

	"github.com/saine1a/stravaanalytics/set"
	"github.com/saine1a/stravaanalytics/stravaaccess"
)

type field struct {
	table string
	name  string
}

func recursiveExplore(typ reflect.Type) (set.Set, []field) {

	fields := make([]field, 0, 100)
	tables := set.Set({make[]string, 0, 100})

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.Kind() == reflect.Struct {
			t, f := recursiveExplore(typ.Field(i).Type)

			tables.AddSet(t)
			fields = append(fields, f...)
		} else {
			tables.Add("dimension")
			fields = append(fields, field{table: typ.Field(i).Tag.Get("dimension"), name: typ.Field(i).Name})
		}
	}

	return tables, fields
}

func (obj *DBaccess) exploreSchema(activities *([]stravaaccess.SummaryActivity)) (set.Set, []field) {

	typ := reflect.TypeOf((*activities)[0])

	return recursiveExplore(typ)
}

func (obj *DBaccess) createSchema(club stravaaccess.Club, activities *([]stravaaccess.SummaryActivity)) {

	// Explore the schema

	tables, fields := obj.exploreSchema(activities)

	for t := range tables {
		fmt.Printf("Table %s\n", tables[t])
	}

	for f := range fields {
		fmt.Printf("Field %s\n", fields[f])
	}

	// First drop the table, in case it already exists

	tableName := fmt.Sprintf("Activity_%s_%d", club.Name, club.ID)

	dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", sqlName(tableName))

	obj.execSQL(dropStmt)

	// Now create the table

	createTableStmt := fmt.Sprintf("CREATE TABLE `%s` (`activityName` VARCHAR(256), `rider` VARCHAR(256), `distance` INT(10))", sqlName(tableName))

	fmt.Println(createTableStmt)

	obj.execSQL(createTableStmt)

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
