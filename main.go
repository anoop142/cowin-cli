package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/cowin"
)

const (
	version = "0.96"
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
	slot := flag.String("slot", "FORENOON", "slot time")
	version := flag.Bool("version", false, "version")

	flag.Parse()
	flag.Usage = func() {
		printAbout()
		help := "Usage :\n"
		help += "\nList :"
		help += "\n  cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b] [-c dd-mm-yyyy]\n"
		help += "  cowin-cli -p pincode \n\n"
		help += "Book Vaccine:"
		help += "\n  cowin-cli -sc -state -d district [-no mobileNumber] [-name Name] [-centers center1,cetner2 ]\n\n"

		fmt.Print(help)
		fmt.Println("Options :")
		flag.PrintDefaults()
	}
	if *pincode != "" || (*state != "" && *district != "") {
		if *schedule {
			cowin.ScheduleVaccine(*state, *district, *pincode,
				*date, *mobileNumber, *name, *centers, *slot, *mAgeLimit)
		} else {
			cowin.PrintCenters(
				*state, *district, *pincode, *vaccine,
				*date, *info, *bookable, *mAgeLimit,
			)
		}
	} else if *version {
		printAbout()
	} else {
		flag.Usage()
	}

}
