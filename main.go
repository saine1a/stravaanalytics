package main

import (
	"fmt"

	"github.com/saine1a/stravaanalytics/db"
	"github.com/saine1a/stravaanalytics/stravaaccess"
)

func main() {

	stravaObj := stravaaccess.Init()

	dbObj := db.Init()

	clubs := stravaObj.GetClubs()

	// Print out

	fmt.Printf("%d clubs found\n", len(*clubs))

	for c := range *clubs {

		activities := stravaObj.GetActivities((*clubs)[c].ID)

		fmt.Printf("Club : %s Activities : %d\n", (*clubs)[c].Name, len(*activities))

		for a := range *activities {
			act := (*activities)[a]
			fmt.Printf("\t%s by %s %s distance %f\n", act.Name, act.Athlete.FirstName, act.Athlete.LastName, act.Distance)
		}

		dbObj.StoreActivities((*clubs)[c], activities)
	}

	return
}
