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

func (center *CentreData) getCenters(districtID string, options Options) {

	u, err := url.Parse(getApiURL(options.Pincode))

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("date", options.Date)

	if options.Vaccine != "" {
		v, ok := vaccinesList[options.Vaccine]

		if !ok {
			log.Fatal("Invalid vaccine")
		}
		q.Add("vaccine", v)
	}

	if options.Pincode != "" {
		q.Add("pincode", options.Pincode)
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

func (center CentreData) printCenterData(options Options) {
	found := false
	for _, v := range center.Centers {
		// skip if the min age limit is greater than specified age
		if options.Age > 0 && options.Age < v.Sessions[0].MinAgeLimit {
			continue
		}
		// skip if  the center is  not bookable
		if options.Bookable {
			totalAvailablity := 0
			for _, vv := range v.Sessions {
				totalAvailablity += vv.AvailableCapacity
			}
			if totalAvailablity < 1 {
				continue
			}
		}
		found = true
		if !options.Info {
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

func PrintCenters(options Options) {
	var center CentreData

	center.getCenters(getDistrictID(options.State, options.District), options)

	center.printCenterData(options)

}
