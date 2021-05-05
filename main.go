package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/states"
)

const (
	version = "0.95"
	author  = "Anoop S"
)

func printAbout() {
	fmt.Printf("cowin-cli v%v %v\n", version, author)
}

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
	centers := flag.String("centers", "", "centers to auto book seperated by ,")
	mAgeLimit := flag.Int("m", 0, "minimum age limit")
	version := flag.Bool("version", false, "version")

	flag.Parse()
	flag.Usage = func() {
		printAbout()
		fmt.Printf("Usage :\n")
		fmt.Println("\nList :")
		fmt.Printf("\n  cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b] [-c dd-mm-yyyy]\n")
		fmt.Printf("  cowin-cli -p pincode \n\n")
		fmt.Println("Book Vaccine:")
		fmt.Printf("\n  cowin-cli -sc -state -d district [-no mobileNumber] [-name Name] [-centers center1,cetner2 ]\n\n")
		fmt.Println("Options :")
		flag.PrintDefaults()
	}
	if *pincode != "" || (*state != "" && *district != "") {
		if *schedule {
			scheduleVaccine(states.GetDistrictID(*state, *district), *pincode,
				*date, *mobileNumber, *name, *centers)
		} else {
			printCenters(states.GetDistrictID(
				*state, *district), *pincode, *vaccine,
				*date, *info, *bookable, *mAgeLimit,
			)
		}
	} else if *version {
		printAbout()
	} else {
		flag.Usage()
	}

}
