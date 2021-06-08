package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type CenterBookable struct {
	Name              string
	Freetype          string
	SessionID         string
	MinAgeLimit       int
	Date              string
	Vaccine           string
	AvailableCapacity int
	DoseType          string
	Slots             []string
}

type beneficariesData struct {
	Beneficiaries []struct {
		BeneficiaryReferenceID string `json:"beneficiary_reference_id"`
		Name                   string `json:"name"`
		Dose1Date              string `json:"dose1_date"`
	} `json:"beneficiaries"`
}

type ScheduleData struct {
	slot               string
	txnId              string
	bearerToken        string
	beneficariesRefIDs []string
	dose               int
	sessionID          string
}

type BadRequest struct {
	Errorcode string `json:"errorCode"`
	Error     string `json:"error"`
}

func getBeneficaries(bearerToken string) (responseCode int, b beneficariesData) {
	auth := true
	resp, statusCode := getReqAuth(beneficiariesURL, bearerToken, auth)

	if statusCode == 200 {
		json.Unmarshal(resp, &b)
	}

	return statusCode, b

}

func getUserSelection(message string, limit int, all bool) int {
	var opt int
	again := false
	fmt.Println()
	for {
		if again {
			fmt.Println("Wrong selection")
		}
		fmt.Print(message)
		fmt.Scanf("%d\n", &opt)

		if opt <= limit || (all && opt == limit+1) {
			break
		} else {
			again = true
		}
	}
	return opt
}

func getDoseNo(doseDate string) int {
	if doseDate == "" {
		return 1
	}
	return 2
}

// getAllbId gets all ref id and a common dose date
func (scheduleData *ScheduleData) getAllbID(b beneficariesData) {
	for _, v := range b.Beneficiaries {
		scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, v.BeneficiaryReferenceID)
	}
	scheduleData.dose = getDoseNo(b.Beneficiaries[0].Dose1Date)
}

func printBeneficaries(b beneficariesData) {
	var all int
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name"})

	for i, v := range b.Beneficiaries {
		table.Append([]string{fmt.Sprint(i), v.Name})
		all = i
	}
	table.Append([]string{fmt.Sprint(all + 1), "All"})

	table.Render()

}

// getBeneficariesID gets list of beneficaries id and dose
func (scheduleData *ScheduleData) getBeneficariesID(b beneficariesData, names string) {

	IDtotalCount := len(b.Beneficiaries)
	if IDtotalCount == 1 {
		scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, b.Beneficiaries[0].BeneficiaryReferenceID)
		scheduleData.dose = getDoseNo(b.Beneficiaries[0].Dose1Date)
		// name specified
	} else if names != "" {
		// get all beneficaries
		if names == "all" {
			scheduleData.getAllbID(b)
		} else {
			nameList := strings.Split(names, ",")
			for _, name := range nameList {
				for _, v := range b.Beneficiaries {
					if strings.EqualFold(v.Name, strings.TrimSpace(name)) {
						scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, v.BeneficiaryReferenceID)
						scheduleData.dose = getDoseNo(v.Dose1Date)
						break
					}
				}
			}

		}

	}
	if len(scheduleData.beneficariesRefIDs) == 0 {
		printBeneficaries(b)
		fmt.Println("use ',' to seperate multiple id")
		for {
			var opt string
			//print beneficaries and prompt user
			fmt.Print("\nEnter Name ID: ")
			fmt.Scanf("%s\n", &opt)

			userIDs := strings.Split(opt, ",")

			for _, v := range userIDs {
				id, _ := strconv.Atoi(v)
				// all mode
				if id == IDtotalCount {
					scheduleData.beneficariesRefIDs = nil
					scheduleData.getAllbID(b)
					break
				}
				if id < IDtotalCount {
					scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, b.Beneficiaries[id].BeneficiaryReferenceID)
					scheduleData.dose = getDoseNo(b.Beneficiaries[id].Dose1Date)
				} else {
					break
				}
			}
			if len(scheduleData.beneficariesRefIDs) != 0 {
				break
			}
		}
	}

}

// printCenterBookable prints centers avaliable for booking
func printCenterBookable(centerList []CenterBookable) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Center", "Free type", "Min Age", "Vaccine", "Dose"})
	for i, v := range centerList {
		table.Append([]string{fmt.Sprint(i), v.Name, v.Freetype, fmt.Sprint(v.MinAgeLimit), v.Vaccine, v.DoseType})
	}
	table.Render()
}
func getSpecifiedCenterSessionID(centerBookable []CenterBookable, specifiedCenters string) CenterBookable {
	var selectedCenter CenterBookable
	if specifiedCenters == "any" {
		// get first session id
		selectedCenter = centerBookable[0]
	} else {
		specifiedCentersList := strings.Split(specifiedCenters, ",")
		for _, specifiedCenter := range specifiedCentersList {
			// remove leading and trailing spaces
			specifiedCenter = strings.TrimSpace(specifiedCenter)
			for _, center := range centerBookable {
				if strings.EqualFold(center.Name, specifiedCenter) {
					selectedCenter = center
				}
			}
		}
	}
	return selectedCenter
}

// getCenterBookable gets centers that are only avaliable for booking
func getCenterBookable(options Options, bearerToken string) []CenterBookable {
	var center CentreData
	var centerBookable []CenterBookable

	if options.Vaccine != "" {
		checkVaccineKnown(options.Vaccine)
	}

	center.getCenters(options, bearerToken)

	for _, v := range center.Centers {
		for _, vv := range v.Sessions {
			doseType := getDoseType(vv.AvailableCapacityDose1, vv.AvailableCapacityDose2)
			/*This is becoming a clusterfuck  */
			if (!options.Bookable || vv.AvailableCapacity >= options.Mslot) && (options.Age == 0 || options.Age >= vv.MinAgeLimit) &&
				(options.Vaccine == "" || checkVaccine(options.Vaccine, vv.Vaccine)) &&
				(options.Dose == 0 || checkDoseType(doseType, options.Dose)) &&
				(options.FreeType == "" || strings.EqualFold(options.FreeType, v.FeeType)) {

				centerBookable = append(centerBookable, CenterBookable{
					Name:              v.Name,
					Freetype:          v.FeeType,
					SessionID:         vv.SessionID,
					Vaccine:           vv.Vaccine,
					MinAgeLimit:       vv.MinAgeLimit,
					Date:              vv.Date,
					AvailableCapacity: vv.AvailableCapacity,
					DoseType:          doseType,
					Slots:             vv.Slots,
				})

			}
		}
	}
	return centerBookable

}

// getSessionID gets session ID and generates OTP
func (scheduleData *ScheduleData) getSessionID(options Options) {

	var opt int
	var selectedCenter CenterBookable
	centerBookable := getCenterBookable(options, scheduleData.bearerToken)

	if len(centerBookable) > 0 {
		// generate OTP only if there is bookable centers && invalid token

		if options.Centers != "" {
			selectedCenter = getSpecifiedCenterSessionID(centerBookable, options.Centers)
		}
		// set session id
		scheduleData.sessionID = selectedCenter.SessionID

		if scheduleData.sessionID == "" {
			printCenterBookable(centerBookable)
			opt = getUserSelection("Enter Center ID :", len(centerBookable)-1, false)

			selectedCenter = centerBookable[opt]
			scheduleData.sessionID = selectedCenter.SessionID
		}
		// Set slot
		if options.Slot == "" {
			scheduleData.slot = selectedCenter.Slots[0]
		} else {
			scheduleData.slot = options.Slot
		}

		fmt.Printf("Center: %v %v Dose-%v %v\n\n", selectedCenter.Name, selectedCenter.Vaccine, selectedCenter.DoseType, selectedCenter.Date)
	} else {
		log.Fatalln("No Centers available for booking")
	}

}

func (scheduleData ScheduleData) scheduleVaccineNow() ([]byte, int) {
	postData := map[string]interface{}{
		"dose":          scheduleData.dose,
		"session_id":    scheduleData.sessionID,
		"slot":          scheduleData.slot,
		"beneficiaries": scheduleData.beneficariesRefIDs,
	}

	jsonBytes, _ := json.Marshal(postData)

	return postReq(appointmentSchedule, jsonBytes, scheduleData.bearerToken)

}

func ScheduleVaccine(options Options) {
	var scheduleData ScheduleData
	var badRequest BadRequest
	var beneficaries beneficariesData
	var OTP, lastRecievedTime, recievedTime string
	var tokenValid = false
	var respCode int
	options.Bookable = true

	if runtime.GOOS == "android" && options.Aotp {
		_, lastRecievedTime = catchOTP()
	}

	if !options.Ntok {
		var ok bool
		scheduleData.bearerToken, ok = loadTokenFromFile(options.TokenFile)
		if ok {
			respCode, beneficaries = getBeneficaries(scheduleData.bearerToken)
			if respCode == 200 {
				tokenValid = true
			}
		}

	}

	if !tokenValid {
		scheduleData.txnId = genOTP(options.MobileNumber)

		if runtime.GOOS == "android" && options.Aotp && checkTermuxAPI() {
			for {
				fmt.Println("Waiting for OTP..")
				OTP, recievedTime = catchOTP()
				if recievedTime != lastRecievedTime {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}

		}
		if OTP == "" {
			OTP = getOTPprompt()
		}

		respCode = scheduleData.validateOTP(OTP)
		// ask 3 times if otp is incorrect
		for i := 0; respCode != 200 && i < 3; i++ {
			fmt.Println("Incorrect OTP")
			respCode = scheduleData.validateOTP(getOTPprompt())

		}

		// write token to file
		if respCode == 200 {
			writeTokenToFile(scheduleData.bearerToken, options.TokenFile)
		}

		respCode, beneficaries = getBeneficaries(scheduleData.bearerToken)
		if respCode != 200 {
			log.Fatalln("Cannot get beneficaries")
		}
	}

	scheduleData.getSessionID(options)

	scheduleData.getBeneficariesID(beneficaries, options.Names)

	resp, statusCode := scheduleData.scheduleVaccineNow()

	switch statusCode {
	case 200:
		fmt.Println("Appointment scheduled successfully!")
		os.Exit(0)
	case 400:
		json.Unmarshal(resp, &badRequest)
		log.Fatalln(badRequest.Error)

	case 401:
		log.Fatalln("Unauthenticated Access")
	case 409:
		log.Fatalln("This vaccination center is completely booked for the selected date")
	case 500:
		log.Fatalln("Internal Server error")
	default:
		log.Fatalln("Error ", statusCode)
	}

}
