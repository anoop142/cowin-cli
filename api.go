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

func getDataDistricts(districtID string, date string) {
	u, err := url.Parse(apiDistrictURL)

	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("district_id", districtID)
	q.Add("date", date)

	u.RawQuery = q.Encode()

	json.Unmarshal(getReq(u.String()), &centreData)

}

func getDataByPincode(pincode string, date string) {
	u, err := url.Parse(apiPincodeURL)

	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("date", date)
	q.Add("pincode", pincode)
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

func printCentersByDistrict(districtID string, date string, printInfo bool) {

	getDataDistricts(districtID, date)

	printCenterData(printInfo)

}

func printCentersByPincode(pincode string, date string, printInfo bool) {

	getDataByPincode(pincode, date)

	printCenterData(printInfo)
}
