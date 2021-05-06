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
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

func getApiURL(pincode string) string {
	if pincode != "" {
		return apiPincodeURL
	} else {
		return apiDistrictURL
	}
}

func (center *CentreData) getCenters(districtID string, pincode string, vaccine string, date string) {

	u, err := url.Parse(getApiURL(pincode))

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("date", date)

	if vaccine != "" {
		v, ok := vaccinesList[vaccine]

		if !ok {
			log.Fatal("Invalid vaccine")
		}
		q.Add("vaccine", v)
	}

	if pincode != "" {
		q.Add("pincode", pincode)
	} else {
		q.Add("district_id", districtID)
	}

	u.RawQuery = q.Encode()
	resp, statusCode := getReqAuth(u.String(), "")

	if statusCode != 200 {
		log.Fatalln(string(resp))
	}

	json.Unmarshal(resp, center)

}

func (center CentreData) printCenterData(printInfo, bookable bool, spAge int) {
	found := false
	for _, v := range center.Centers {
		// skip if the min age limit is greater than specified age
		if spAge > 0 && spAge < v.Sessions[0].MinAgeLimit {
			continue
		}
		// skip if  the center is  not bookable
		if bookable {
			totalAvailablity := 0
			for _, vv := range v.Sessions {
				totalAvailablity += vv.AvailableCapacity
			}
			if totalAvailablity < 1 {
				continue
			}
		}
		found = true
		if !printInfo {
			fmt.Printf("%v ", v.Name)
			if v.FeeType != "Free" {
				fmt.Printf("Paid")
			}
			fmt.Println()
		} else {
			for _, vv := range v.Sessions {
				fmt.Printf("%v  %v  %v  %v  %v %v+\n", v.Name, v.FeeType, vv.Date, vv.AvailableCapacity, vv.Vaccine, vv.MinAgeLimit)
			}
		}
	}
	if !found {
		os.Exit(1)
	}
}

func PrintCenters(state, district, pincode, vaccine, date string,
	printInfo, bookable bool, spAge int) {
	var center CentreData

	center.getCenters(getDistrictID(state, district), pincode, vaccine, date)

	center.printCenterData(printInfo, bookable, spAge)

}
