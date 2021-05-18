package cowin

import (
	"log"
	"strings"
)

var vaccinesList = map[string]bool{
	"COVAXIN":    true,
	"COVISHIELD": true,
	"SPUTNIK V":  true,
}

func getVaccineName(vaccine string) string {
	if vaccine != "" {
		vaccine = strings.ToUpper(vaccine)
		_, ok := vaccinesList[vaccine]
		if !ok {
			log.Fatalln("Invalid vaccine!")
		}
	}

	return vaccine
}
