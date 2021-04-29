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
	pincode := flag.String("pin", "", "pincode")
	state := flag.String("state", "", "state")
	district := flag.String("district", "", "district")
	date := flag.String("date", getDate(), "date dd-mm-yyyy")
	vaccine := flag.String("vaccine", "", "vaccine name")
	info := flag.Bool("info", false, "full info")

	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("Usage :\n")
		fmt.Printf("\n  cowin-cli --state state --district district [--vaccine vaccine name] [--info] [--date dd-mm-yyyy]\n")
		fmt.Printf("  cowin-cli --pin pincode \n\n")
		flag.PrintDefaults()
	}
	if *pincode != "" || (*state != "" && *district != "") {
		printCenters(states.GetDistrictID(*state, *district), *pincode, *vaccine, *date, *info)
	} else {
		flag.Usage()
	}

}
