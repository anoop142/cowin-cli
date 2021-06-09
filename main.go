package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/anoop142/cowin-cli/cowin"
)

const (
	version = "1.7.9"
	author  = "Anoop S"
)

func printAbout() {
	fmt.Printf("cowin-cli v%v %v\n", version, author)
}

// get tomorrow's date
func getDate(noCacheBypass bool) string {
	dateNow := time.Now()
	dateTommorrow := dateNow.AddDate(0, 0, 1)

	if noCacheBypass {
		return dateTommorrow.Format("02-01-2006")
	} else {
		// hack to bypass cache
		return dateTommorrow.Format("02-01-06")
	}
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
	token := flag.String("token", "token.txt", "file to read/write token")
	freeType := flag.String("t", "", "free type")
	slot := flag.String("slot", "", "slot time")
	minSlot := flag.Int("ms", 1, "minimum slots")
	protected := flag.Bool("p", false, "use protected URL to list")
	noCacheBypass := flag.Bool("ncb", false, "don't use cache bypass")
	version := flag.Bool("version", false, "version")
	help := flag.Bool("help", false, "help")

	flag.Parse()
	flag.Usage = func() {
		printAbout()
		helpMsg := "Usage :\n"
		helpMsg += "\nList :"
		helpMsg += "\n  cowin-cli -s state -d district [-v vaccine] [-m age] [-i] [-b] [-c dd-mm-yyyy][-dose dose] [-t freeType] [-ms minimumSlot] [-p] [-ncb]\n\n"
		helpMsg += "Book Vaccine:"
		helpMsg += "\n  cowin-cli -sc -s state -d district [-no mobileNumber] [-v vaccine] [-m age] [-names name1,name2] [-centers center1,cetner2 ] [-slot slotTime] [-aotp] [-ntok] [-token tokenFile] [-dose dose] [-c dd-mm-yyyy] [-t freeType] [-ms minimumSlot] [-ncb] \n\n"
		helpMsg += "Generate Token:"
		helpMsg += "\n	cowin-cli -gen [-no mobileNumber] [-token tokenFile]  \n\n"
		fmt.Print(helpMsg)
	}
	if *state != "" && *district != "" {
		// set date if not specified
		if *date == "" {
			*date = getDate(*noCacheBypass)
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
			Protected:    *protected,
			Mslot:        *minSlot,
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
		fmt.Println("Options :")
		flag.PrintDefaults()
	} else {
		flag.Usage()
	}

}
