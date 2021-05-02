package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
)

// genOTP generates OTP and return txnId
func (scheduleData *ScheduleData) genOTP(mobileNumber string) {

	if mobileNumber == "" {
		fmt.Printf("\nEnter Mobile Number : ")
		fmt.Scanf("%s\n", &mobileNumber)
		fmt.Println()
	}

	var otpResp map[string]interface{}

	postData := map[string]string{
		"secret": "U2FsdGVkX1/kJeCrPtUR3tpExWzWcm21DTIs03CdaVIWLrKmWI9HAnMZ91KVhzZoMGlIKi/83Iy4khgzhR107A==",
		"mobile": mobileNumber,
	}

	jsonBytes, _ := json.Marshal(postData)

	resp, statusCode := postReq(otpGenURL, jsonBytes, "")

	if statusCode != 200 {
		log.Fatalln("otp failed")
	}
	json.Unmarshal(resp, &otpResp)

	txnId, ok := otpResp["txnId"]

	if !ok {
		log.Fatalln("cannot get txnId")
	}

	scheduleData.txnId = fmt.Sprintf("%v", txnId)

}

// getOTP prompts user for OTP
func getOTPprompt() string {
	var otp string
	fmt.Print("Enter OTP :")
	fmt.Scanf("%s\n", &otp)
	return otp
}

// validateOTP validates OTP ans gets bearer token
func (scheduleData *ScheduleData) validateOTP(otp string) {
	otpSha256 := sha256.Sum256([]byte(otp))

	var loginData map[string]interface{}

	postData := map[string]string{
		"otp":   fmt.Sprintf("%x", otpSha256),
		"txnId": scheduleData.txnId,
	}
	jsonBytes, _ := json.Marshal(postData)
	resp, statusCode := postReq(otpValURL, jsonBytes, "")

	if statusCode == 200 {

		json.Unmarshal(resp, &loginData)

		bearerToken, ok := loginData["token"]

		if !ok {
			log.Fatalln("Cannot get token")
		}

		scheduleData.bearerToken = fmt.Sprintf("%v", bearerToken)
	}

}
