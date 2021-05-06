package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type StatesData struct {
	States []struct {
		StateID   int    `json:"state_id"`
		StateName string `json:"state_name"`
	} `json:"states"`
}

type DistrictsData struct {
	Districts []struct {
		DistrictID   int    `json:"district_id"`
		DistrictName string `json:"district_name"`
	} `json:"districts"`
}

func getStateID(state string) int {
	var statesData StatesData
	var stateID int
	state = strings.Title(state)

	resp, statusCode := getReqAuth(statesURL, "")

	if statusCode != 200 {
		log.Fatalln(string(resp))
	}

	json.Unmarshal(resp, &statesData)

	for _, v := range statesData.States {
		if v.StateName == state {
			stateID = v.StateID
			break
		}
	}

	if stateID == 0 {
		log.Fatalln("Invalid state!")
	}

	return stateID

}

func getDistrictID(state, district string) string {
	var districtsData DistrictsData
	var districtID int

	district = strings.Title(district)

	stateID := getStateID(state)

	resp, statusCode := getReqAuth(fmt.Sprintf("%v/%v", districtsURL, stateID), "")

	if statusCode != 200 {
		log.Fatalln(string(resp))
	}

	json.Unmarshal(resp, &districtsData)

	for _, v := range districtsData.Districts {
		if v.DistrictName == district {
			districtID = v.DistrictID
		}
	}

	if districtID == 0 {
		log.Fatalln("Invalid district!")
	}

	return fmt.Sprint(districtID)

}
