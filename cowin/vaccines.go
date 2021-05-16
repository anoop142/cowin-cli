package cowin

import "log"

var vaccinesList = map[string]string{
	"covaxin":    "COVAXIN",
	"covishield": "COVISHIELD",
}

func getVaccineName(vaccine string) string {
	var vaccineName string
	var ok bool
	if vaccine != "" {
		vaccineName, ok = vaccinesList[vaccine]
		if !ok {
			log.Fatalln("Invalid vaccine")
		}
	}
	return vaccineName
}
