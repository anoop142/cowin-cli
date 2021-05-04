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
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

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

	json.Unmarshal(getReq(u.String()), center)

}

func (center CentreData) printCenterData(printInfo,
	bookable bool, mAgeLimit int) {
	for _, v := range center.Centers {
		// skip if  the center is  not bookable
		if bookable && v.Sessions[0].AvailableCapacity < 1 {
			continue
		}
		// skip if the specified age limit is less than min age limit
		if mAgeLimit > 0 && mAgeLimit < v.Sessions[0].MinAgeLimit {
			continue
		}

		fmt.Printf("%v ", v.Name)
		if v.FeeType != "Free" {
			fmt.Printf("Paid")
		}
		fmt.Println()

		if printInfo {
			fmt.Printf("  %v - %v  %v %v+\n", v.Sessions[0].Date, v.Sessions[0].AvailableCapacity, v.Sessions[0].Vaccine, v.Sessions[0].MinAgeLimit)

		}
	}
}

func printCenters(districtID, pincode, vaccine, date string,
	printInfo, bookable bool, mAgeLimit int) {
	var center CentreData

	center.getCenters(districtID, pincode, vaccine, date)

	center.printCenterData(printInfo, bookable, mAgeLimit)

}
