package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CentreData struct {
	Centers []struct {
		Name     string `json:"name"`
		FeeType  string `json:"fee_type"`
		Sessions []struct {
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

var centreData CentreData

func getReq(URL string) []byte {
	resp, err := http.Get(URL)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body

}

func getApiURL(pincode string) string {
	if pincode != "" {
		return apiPincodeURL
	} else {
		return apiDistrictURL
	}
}

func getCenters(districtID string, pincode string, vaccine string, date string) {

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

	json.Unmarshal(getReq(u.String()), &centreData)

}

func printCenterData(printInfo bool, bookable bool) {
	for _, v := range centreData.Centers {

		// skip if  the center is  not bookable
		if bookable {
			totalCapacity := 0
			for _, vv := range v.Sessions {
				totalCapacity += vv.AvailableCapacity
			}
			if totalCapacity < 1 {
				continue
			}
		}

		fmt.Printf("%v ", v.Name)
		if v.FeeType != "Free" {
			fmt.Printf("Paid")
		}
		fmt.Println()

		if printInfo {
			for _, vv := range v.Sessions {
				fmt.Printf("  %v - %v  %v %v+\n", vv.Date, vv.AvailableCapacity, vv.Vaccine, vv.MinAgeLimit)
			}
		}
	}
}

func printCenters(districtID, pincode, vaccine, date string, printInfo bool, bookable bool) {

	getCenters(districtID, pincode, vaccine, date)

	printCenterData(printInfo, bookable)

}
