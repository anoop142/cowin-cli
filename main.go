package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/anoop142/cowin-cli/states"
)

func getDate() string {
	dateNow := time.Now()
	dateTommorrow := dateNow.AddDate(0, 0, 1)
	return dateTommorrow.Format("02-01-2006")
}

func main() {
	state := flag.String("state", "", "state")
	district := flag.String("district", "", "district")
	date := flag.String("date", getDate(), "date dd-mm-yyyy")
	info := flag.Bool("info", false, "full info")

	flag.Parse()
	flag.Usage = func() {
		fmt.Printf("%s --state state --district district [--info] [--date dd-mm-yyyy]\n", os.Args[0])
		flag.PrintDefaults()
	}

	if *state != "" && *district != "" {
		apiURL := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/calendarByDistrict?district_id=%s&date=%s", states.GetDistrictID(*state, *district), *date)
		printCenters(apiURL, *info)
	} else {
		flag.Usage()
	}

}
