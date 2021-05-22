package cowin

import (
	"fmt"
	"log"
	"os"
)

func writeTokenToFile(token, tokenFile string) bool {

	err := os.WriteFile(tokenFile, []byte(token), 0644)

	return err == nil
}

func loadTokenFromFile(tokenFile string) (string, bool) {
	tokenBytes, err := os.ReadFile(tokenFile)

	return string(tokenBytes), err == nil

}

// Generates token file
func GenerateToken(number, tokenFile string) {
	var scheduleData ScheduleData

	scheduleData.txnId = genOTP(number)
	if scheduleData.validateOTP(getOTPprompt()) == 200 {
		if writeTokenToFile(scheduleData.bearerToken, tokenFile) {
			fmt.Println("Written to", tokenFile)
		}

	} else {
		log.Fatalln("Cannot get token")
	}

}
