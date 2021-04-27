package states

import "log"

var klDistricts = map[string]string{
	"alappuzha":          "301",
	"ernakulam":          "307",
	"idukki":             "306",
	"kannur":             "297",
	"kasaragod":          "295",
	"kollam":             "298",
	"kottayam":           "304",
	"kozhikode":          "305",
	"malappuram":         "302",
	"palakkad":           "308",
	"pathanamthitta":     "300",
	"thiruvananthapuram": "296",
	"thrissur":           "303",
	"wayanad":            "299",
}

func getDistrictKL(name string) string {
	id, ok := klDistricts[name]
	if !ok {
		log.Fatalln("Invalid district")
	}
	return id
}
