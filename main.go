package main

import (
	"fmt"

	"github.com/saine1a/stravaanalytics/stravaaccess"
)

func main() {

	stravaObj := stravaaccess.Init()

	clubs := stravaObj.GetClubs()

	// Print out

	fmt.Printf("%d clubs found\n", len(*clubs))
	for c := range *clubs {
		fmt.Printf("Club : %s\n", (*clubs)[c].Name)
	}

}
