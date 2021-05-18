package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/cowin"
)

const (
	version = "1.4.5"
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
	vaccine := flag.String("v", "", "vaccines separated by ,")
	info := flag.Bool("i", false, "full info")
	bookable := flag.Bool("b", false, "bookable only")
	schedule := flag.Bool("sc", false, "schedule vaccine ")
	mobileNumber := flag.String("no", "", "mobile number")
	name := flag.String("name", "", "registered name")
	centers := flag.String("centers", "", "centers to auto book separated by ,")
	age := flag.Int("m", 0, "minimum age limit")
	dose := flag.Int("dose", 0, "dose type")
	aotp := flag.Bool("aotp", false, "auto capture otp for termux")
	ntok := flag.Bool("ntok", false, "don't reuse token")
	slot := flag.String("slot", "FORENOON", "slot time")
	version := flag.Bool("version", false, "version")

	flag.Parse()
	flag.Usage = func() {
		printAbout()
		help := "Usage :\n"
		help += "\nList :"
		help += "\n  cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b] [-c dd-mm-yyyy][-dose dose]\n"
		help += "  cowin-cli -p pincode \n\n"
		help += "Book Vaccine:"
		help += "\n  cowin-cli -sc -state -d district [-no mobileNumber] [-v vaccine] [-m age] [-name Name] [-centers center1,cetner2 ] [-slot slotTime] [-aotp] [-ntok] [-dose dose]\n\n"

		fmt.Print(help)
		fmt.Println("Options :")
		flag.PrintDefaults()
	}
	if *pincode != "" || (*state != "" && *district != "") {

		options := cowin.Options{
			Pincode:      *pincode,
			State:        *state,
			District:     *district,
			Date:         *date,
			Vaccine:      *vaccine,
			Info:         *info,
			Bookable:     *bookable,
			Schedule:     *schedule,
			MobileNumber: *mobileNumber,
			Name:         *name,
			Centers:      *centers,
			Age:          *age,
			Slot:         *slot,
			Aotp:         *aotp,
			Ntok:         *ntok,
			Dose:         *dose,
		}
		if *schedule {
			cowin.ScheduleVaccine(options)
		} else {
			cowin.PrintCenters(options)
		}
	} else if *version {
		printAbout()
	} else {
		flag.Usage()
	}

}
