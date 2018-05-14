package db

import (
	"fmt"
	"reflect"

	"github.com/saine1a/stravaanalytics/stravaaccess"
	"github.com/saine1a/stravaanalytics/utils"
)

type field struct {
	name     string
	dataType string
}

type visitor func(field reflect.StructField)

func recursivelyScan(typ reflect.Type, visitorFunc visitor) {

	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Type.Kind() == reflect.Struct {
			recursivelyScan(typ.Field(i).Type, visitorFunc)
		} else {
			visitorFunc(typ.Field(i))
		}
	}
}

func (obj *DBaccess) scanActivities(activities *([]stravaaccess.SummaryActivity), visitorFunc visitor) {

	typ := reflect.TypeOf((*activities)[0])

	recursivelyScan(typ, visitorFunc)
}

func (obj *DBaccess) exploreSchema(activities *([]stravaaccess.SummaryActivity)) *utils.HierarchicalSet {

	schema := utils.CreateHierarchicalSet()

	createSchemaVisitor := func(theField reflect.StructField) {
		theType := theField.Tag.Get("type")
		theDimension := theField.Tag.Get("dimension")
		schema.Add(theDimension, field{name: theField.Name, dataType: theType})
	}

	obj.scanActivities(activities, createSchemaVisitor)

	return schema
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

		fieldType := fields[f].(field).dataType
		if fieldType == "" {
			fieldType = "VARCHAR(256)"
		}
		fieldExpr = fieldExpr + fmt.Sprintf("`%s` %s", fields[f].(field).name, fieldType)
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

	// New insert the values

	/*
		for a := range *activities {

			activity := (*activities)[a]

			for t := range tableSlice {

			}



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
