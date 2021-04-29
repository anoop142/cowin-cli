package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CentreData struct {
	Centers []struct {
		Name     string `json:"name"`
		FeeType  string `json:"fee_type"`
		Sessions []struct {
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
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

func getCenters(districtID string, pincode string, vaccine string, date string) {
	u, err := url.Parse(apiDistrictURL)

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("date", date)

	if vaccine != "" {
		q.Add("vaccine", strings.ToUpper(vaccine))
	}

	if pincode != "" {
		q.Add("pincode", pincode)
	} else {
		q.Add("district_id", districtID)
	}

	u.RawQuery = q.Encode()

	json.Unmarshal(getReq(u.String()), &centreData)

}

func printCenterData(printInfo bool) {
	for _, v := range centreData.Centers {
		fmt.Printf("%v ", v.Name)
		if v.FeeType != "Free" {
			fmt.Printf("Paid")
		}
		fmt.Println()

		if printInfo {
			for _, vv := range v.Sessions {
				fmt.Printf("  %v - %v\n", vv.Date, vv.AvailableCapacity)
			}
		}
	}
}

func printCenters(districtID, pincode, vaccine, date string, printInfo bool) {

	getCenters(districtID, pincode, vaccine, date)

	printCenterData(printInfo)

}
