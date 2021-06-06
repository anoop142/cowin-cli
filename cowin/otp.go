package cowin

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

const termuxSmsApi = "termux-sms-list"

// genOTP generates OTP and return txnId
func genOTP(mobileNumber string) string {
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

	return fmt.Sprintf("%v", txnId)

}

// getOTP prompts user for OTP
func getOTPprompt() string {
	var otp string
	fmt.Print("Enter OTP :")
	fmt.Scanf("%s\n", &otp)
	return otp
}

// validateOTP validates OTP ans gets bearer token
func (scheduleData *ScheduleData) validateOTP(otp string) (statusCode int) {
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
	return statusCode

}

func checkTermuxAPI() bool {
	_, err := exec.LookPath(termuxSmsApi)

	return err == nil
}

// Catch otp for termux
func catchOTP() (string, string) {
	// Extraction needs refactoring
	const otpSubstringLeft = `Your OTP to register/access CoWIN is`
	const otpSubstringRight = `It will be valid for 3 minutes. - CoWIN`
	const timeSubstringLeft = `"received":`
	const timeSubstringRight = `,`

	out, _ := exec.Command(termuxSmsApi, "-l", "1").Output()

	msgList := string(out)

	r := regexp.MustCompile(otpSubstringLeft + `(?)(.*)` + otpSubstringRight)
	OTP := r.FindString(msgList)
	OTP = strings.TrimLeft(strings.TrimRight(OTP, otpSubstringRight), otpSubstringLeft)
	r = regexp.MustCompile(timeSubstringLeft + `(?)(.*)` + timeSubstringRight)
	recievedTime := r.FindString(msgList)
	recievedTime = strings.TrimLeft(strings.TrimRight(recievedTime, timeSubstringRight), timeSubstringLeft)

	return OTP, recievedTime

}
