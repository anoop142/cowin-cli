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

// checks if the vaccine is in the known vaccine list
func checkVaccineKnown(vaccines string) {
	vaccinesSupplied := strings.Split(vaccines, ",")
	for _, v := range vaccinesSupplied {
		v = strings.ToUpper(strings.TrimSpace(v))
		_, ok := vaccinesList[v]
		if !ok {
			log.Fatalln("Invalid vaccine: ", v)
		}

	}
}

// checks the available vaccine is int the supplied list
func checkVaccine(suppliedVaccines, availVaccine string) bool {
	ok := false
	v := strings.Split(suppliedVaccines, ",")
	for _, vaccine := range v {
		vaccine = strings.TrimSpace(vaccine)
		if strings.EqualFold(vaccine, availVaccine) {
			ok = true
			break
		}

	}
	return ok
}
