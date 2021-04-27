package states

import "log"

func GetDistrictID(state string, district string) string {
	var id string
	var ok bool

	switch state {
	case "kerala":
		id, ok = kerala[district]
	default:
		log.Fatalln("Invalid state")
	}

	if !ok {
		log.Fatalln("Invalid district")
	}

	return id
}
