package cowin

import "log"

var vaccinesList = map[string]string{
	"covaxin":    "COVAXIN",
	"covishield": "COVISHIELD",
}

func GetVaccineName(vaccine string) string {
	vaccineName, ok := vaccinesList[vaccine]
	if !ok {
		log.Fatalln("Invalid vaccine")
	}
	return vaccineName
}
