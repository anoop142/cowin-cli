package cowin

import (
	"os"
)

const tokenFile = "token.txt"

func writeTokenToFile(token string) bool {

	err := os.WriteFile(tokenFile, []byte(token), 0644)

	return err == nil
}

func loadTokenFromFile() (string, bool) {
	tokenBytes, err := os.ReadFile(tokenFile)

	return string(tokenBytes), err == nil

}
