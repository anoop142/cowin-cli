package states

import "log"

func GetDistrictID(state string, district string) string {
	switch state {
	case "kerala":
		return getDistrictKL(district)
	default:
		log.Fatalln("Invalid state")
	}

	return ""
}
