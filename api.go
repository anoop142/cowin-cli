package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CentreData struct {
	Centers []struct {
		Name         string `json:"name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Lat          int    `json:"lat"`
		Long         int    `json:"long"`
		From         string `json:"from"`
		To           string `json:"to"`
		FeeType      string `json:"fee_type"`
		Sessions     []struct {
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

var centreData CentreData

func getApiData(apiURL string) {
	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &centreData)

}

func printCenters(apiURL string, printInfo bool) {
	getApiData(apiURL)
	for _, v := range centreData.Centers {
		fmt.Printf("%v\n", v.Name)
		if printInfo {
			for _, vv := range v.Sessions {
				fmt.Printf("  %v - %v\n", vv.Date, vv.AvailableCapacity)
			}
		}
	}
}
