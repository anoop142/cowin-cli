package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/cowin"
)

const (
	version = "1.6.0"
	author  = "Anoop S"
)

func printAbout() {
	fmt.Printf("cowin-cli v%v %v\n", version, author)
}

// get tomorrow's date
func getDate() string {
	dateNow := time.Now()
	dateTommorrow := dateNow.AddDate(0, 0, 1)
	return dateTommorrow.Format("02-01-2006")
}

func main() {
	state := flag.String("s", "", "state")
	district := flag.String("d", "", "district")
	date := flag.String("c", "", "date dd-mm-yyyy")
	vaccine := flag.String("v", "", "vaccines separated by ,")
	info := flag.Bool("i", false, "full info")
	bookable := flag.Bool("b", false, "bookable only")
	schedule := flag.Bool("sc", false, "schedule vaccine ")
	gen := flag.Bool("gen", false, "generate token text file.")
	mobileNumber := flag.String("no", "", "mobile number")
	names := flag.String("names", "", "beneficiaries name separated by ,")
	centers := flag.String("centers", "", "centers to auto book separated by ,")
	age := flag.Int("m", 0, "minimum age limit")
	dose := flag.Int("dose", 0, "dose type")
	aotp := flag.Bool("aotp", false, "auto capture otp for termux")
	ntok := flag.Bool("ntok", false, "don't reuse token")
	token := flag.String("token", "token.txt", "file to write token")
	freeType := flag.String("t", "", "free type")
	slot := flag.String("slot", "", "slot time")
	version := flag.Bool("version", false, "version")
	help := flag.Bool("help", false, "help")

	flag.Parse()
	flag.Usage = func() {
		printAbout()
		helpMsg := "Usage :\n"
		helpMsg += "\nList :"
		helpMsg += "\n  cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b] [-c dd-mm-yyyy][-dose dose] [-t freeType]\n\n"
		helpMsg += "Book Vaccine:"
		helpMsg += "\n  cowin-cli -sc -s state -d district [-no mobileNumber] [-v vaccine] [-m age] [-names name1,name2] [-centers center1,cetner2 ] [-slot slotTime] [-aotp] [-ntok] [-token tokenFile] [-dose dose] [-t freeType]\n\n"
		helpMsg += "Generate Token:"
		helpMsg += "\n	cowin-cli -gen [-no mobileNumber] [-token tokenFile]  \n\n"
		fmt.Print(helpMsg)
		fmt.Println("Options :")
		flag.PrintDefaults()
	}
	if *state != "" && *district != "" {
		// set date if not specified
		if *date == "" {
			*date = getDate()
		}

		options := cowin.Options{
			State:        *state,
			District:     *district,
			Date:         *date,
			Vaccine:      *vaccine,
			Info:         *info,
			Bookable:     *bookable,
			Schedule:     *schedule,
			MobileNumber: *mobileNumber,
			Names:        *names,
			Centers:      *centers,
			Age:          *age,
			Slot:         *slot,
			Aotp:         *aotp,
			Ntok:         *ntok,
			Dose:         *dose,
			TokenFile:    *token,
			FreeType:     *freeType,
		}
		if *schedule {
			cowin.ScheduleVaccine(options)
		} else {
			cowin.PrintCenters(options)
		}

	} else if *gen {
		cowin.GenerateToken(*mobileNumber, *token)
	} else if *version {
		printAbout()
	} else if *help {
		flag.Usage()
	} else {
		flag.Usage()
	}

}
