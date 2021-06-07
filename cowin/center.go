package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)

type CentreData struct {
	Centers []struct {
		Name     string `json:"name"`
		FeeType  string `json:"fee_type"`
		Sessions []struct {
			SessionID              string   `json:"session_id"`
			Date                   string   `json:"date"`
			AvailableCapacity      int      `json:"available_capacity"`
			MinAgeLimit            int      `json:"min_age_limit"`
			Vaccine                string   `json:"vaccine"`
			Slots                  []string `json:"slots"`
			AvailableCapacityDose1 int      `json:"available_capacity_dose1"`
			AvailableCapacityDose2 int      `json:"available_capacity_dose2"`
		} `json:"sessions"`
	} `json:"centers"`
}

func (center *CentreData) getCenters(options Options, bearerToken string) {
	var (
		URL  string
		auth = false
	)

	if options.Protected || options.Schedule {
		URL = calenderDistrictURL
		auth = true
	} else {
		URL = calendarDistrictPublicURL
	}
	u, err := url.Parse(URL)

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("date", options.Date)

	districtID := getDistrictID(options.State, options.District)
	q.Add("district_id", districtID)

	u.RawQuery = q.Encode()
	resp, statusCode := getReqAuth(u.String(), bearerToken, auth)

	if statusCode != 200 {
		log.Fatalln(string(resp))
	}

	json.Unmarshal(resp, center)

}
func getDoseType(dose1, dose2 int) string {
	var doseType string
	if dose1 > 1 && dose2 > 1 {
		doseType = "both"
	} else if dose1 > 1 {
		doseType = "1"
	} else if dose2 > 1 {
		doseType = "2"
	} else {
		doseType = "none"
	}
	return doseType
}
func checkDoseType(dosType string, specifiedDose int) bool {
	ok := false
	switch dosType {
	case "both":
		ok = true
	case fmt.Sprint(specifiedDose):
		ok = true
	}
	return ok
}

func PrintCenters(options Options) {
	var (
		bearerToken string
		ok          bool
	)

	if options.Protected {
		bearerToken, ok = loadTokenFromFile(options.TokenFile)
		if !ok {
			log.Fatalln(options.TokenFile, "not found!")
		}
	}
	center := getCenterBookable(options, bearerToken)
	if len(center) > 0 {
		for _, v := range center {
			if options.Info {
				fmt.Printf("%v  %v  %v  %v %v %v Dose-%v\n", v.Name, v.Freetype, v.Date, v.AvailableCapacity, v.Vaccine, v.MinAgeLimit, v.DoseType)
			} else {
				fmt.Printf("%s ", v.Name)
				if v.Freetype != "Free" {
					fmt.Print("Paid")
				}
				fmt.Println()
			}

		}
	} else {
		os.Exit(1)
	}

}
