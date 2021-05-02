package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/states"
)

func getDate() string {
	dateNow := time.Now()
	dateTommorrow := dateNow.AddDate(0, 0, 1)
	return dateTommorrow.Format("02-01-2006")
}

func main() {
	pincode := flag.String("p", "", "pincode")
	state := flag.String("s", "", "state")
	district := flag.String("d", "", "district")
	date := flag.String("c", getDate(), "date dd-mm-yyyy")
	vaccine := flag.String("v", "", "vaccine name")
	info := flag.Bool("i", false, "full info")
	bookable := flag.Bool("b", false, "bookable only")
	schedule := flag.Bool("sc", false, "schedule vaccine ")
	mobileNumber := flag.String("no", "", "mobile number")
	name := flag.String("name", "", "registered name")

	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("Usage :\n")
		fmt.Println("\nList :")
		fmt.Printf("\n  cowin-cli -s state -d district [-v vaccine ] [-i] [-b] [-c dd-mm-yyyy]\n")
		fmt.Printf("  cowin-cli -p pincode \n\n")
		fmt.Println("Book Vaccine:")
		fmt.Printf("\n  cowin-cli -sc -state -d district [-no mobileNumber] [-name Name]\n\n")
		fmt.Println("Options :")
		flag.PrintDefaults()
	}
	if *pincode != "" || (*state != "" && *district != "") {
		if *schedule {
			scheduleVaccine(states.GetDistrictID(*state, *district), *pincode, *date, *mobileNumber, *name)
		} else {
			printCenters(states.GetDistrictID(*state, *district), *pincode, *vaccine, *date, *info, *bookable)
		}
	} else {
		flag.Usage()
	}

}
