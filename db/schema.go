package db

import (
	"fmt"
	"reflect"

	"github.com/saine1a/stravaanalytics/stravaaccess"
	"github.com/saine1a/stravaanalytics/utils"
)

type field struct {
	table string
	name  string
}

func recursiveExplore(typ reflect.Type) *utils.HierarchicalSet {

	schema := utils.CreateHierarchicalSet()

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.Kind() == reflect.Struct {
			subSet := recursiveExplore(typ.Field(i).Type)

			schema.AddHierarchicalSet(subSet)
		} else {
			schema.Add(typ.Field(i).Tag.Get("dimension"), typ.Field(i).Name)
		}
	}

	return schema
}

func (obj *DBaccess) exploreSchema(activities *([]stravaaccess.SummaryActivity)) *utils.HierarchicalSet {

	typ := reflect.TypeOf((*activities)[0])

	return recursiveExplore(typ)
}

func (obj *DBaccess) createTable(club stravaaccess.Club, table string, fieldSet *utils.Set) {
	// First drop the table, in case it already exists

	tableName := fmt.Sprintf("Activity_%s_%d_%s", club.Name, club.ID, table)

	dropStmt := fmt.Sprintf("DROP TABLE IF EXISTS %s", sqlName(tableName))

	obj.execSQL(dropStmt)

	// Determine the fields

	fields := fieldSet.Slice()

	fieldExpr := ""

	for f := range fields {
		if f > 0 {
			fieldExpr = fieldExpr + ","
		}

		fieldExpr = fieldExpr + fmt.Sprintf("`%s` VARCHAR(256)", fields[f].(string))
	}

	// Now create the table

	createTableStmt := fmt.Sprintf("CREATE TABLE `%s` (%s)", sqlName(tableName), fieldExpr)

	fmt.Println(createTableStmt)

	obj.execSQL(createTableStmt)
}

func (obj *DBaccess) createSchema(club stravaaccess.Club, activities *([]stravaaccess.SummaryActivity)) {

	// Explore the schema

	schema := obj.exploreSchema(activities)

	tableSlice := schema.GetKeys()
	for t := range tableSlice {
		obj.createTable(club, tableSlice[t], schema.GetSecondLevelSet(tableSlice[t]))
	}

	/*
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
	*/
}
