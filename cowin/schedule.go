package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
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
	captcha            string
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
	fmt.Println()
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
func (scheduleData *ScheduleData) getBeneficariesID(b beneficariesData, name string) {
	var opt int

	IDtotalCount := len(b.Beneficiaries)
	if IDtotalCount == 1 {
		scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, b.Beneficiaries[0].BeneficiaryReferenceID)
		scheduleData.dose = getDoseNo(b.Beneficiaries[0].Dose1Date)
		// name specified
	} else if name != "" {
		// get all beneficaries
		if name == "all" {
			scheduleData.getAllbID(b)
		} else {
			for _, v := range b.Beneficiaries {
				if strings.EqualFold(v.Name, name) {
					scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, v.BeneficiaryReferenceID)
					scheduleData.dose = getDoseNo(v.Dose1Date)
					break
				}
			}

		}

	}
	if len(scheduleData.beneficariesRefIDs) == 0 {
		//print beneficaries and prompt user
		printBeneficaries(b)
		opt = getUserSelection("Enter name ID : ", IDtotalCount, true)

		// get all beneficaries
		if opt == IDtotalCount {
			scheduleData.getAllbID(b)
			// append chosen one
		} else {
			scheduleData.beneficariesRefIDs = append(scheduleData.beneficariesRefIDs, b.Beneficiaries[opt].BeneficiaryReferenceID)
			scheduleData.dose = getDoseNo(b.Beneficiaries[opt].Dose1Date)
		}
	}

}

// printCenterBookable prints centers avaliable for booking
func printCenterBookable(centerList []CenterBookable) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Center", "Free type", "Min Age", "Vaccine"})
	for i, v := range centerList {
		table.Append([]string{fmt.Sprint(i), v.Name, v.Freetype, fmt.Sprint(v.MinAgeLimit), v.Vaccine})
	}
	table.Render()
}
func getSpecifiedCenterSessionID(centerBookable []CenterBookable, specifiedCenters string) string {
	var sessionId, centerName, vaccine string
	if specifiedCenters == "any" {
		// get first session id
		sessionId = centerBookable[0].SessionID
		vaccine = centerBookable[0].Vaccine
		centerName = centerBookable[0].Name
	} else {
		specifiedCentersList := strings.Split(specifiedCenters, ",")
		for _, specifiedCenter := range specifiedCentersList {
			// remove leading and trailing spaces
			specifiedCenter = strings.TrimSpace(specifiedCenter)
			for _, center := range centerBookable {
				if strings.EqualFold(center.Name, specifiedCenter) {
					sessionId = center.SessionID
					vaccine = center.Vaccine
					centerName = center.Name
				}
			}
		}
	}
	if sessionId != "" {
		fmt.Println("Center: ", centerName, vaccine)
	}
	return sessionId
}

// getCenterBookable gets centers that are only avaliable for booking
func getCenterBookable(options Options) []CenterBookable {
	var center CentreData
	var centerBookable []CenterBookable

	if options.Vaccine != "" {
		checkVaccineKnown(options.Vaccine)
	}

	center.getCenters(options)

	for _, v := range center.Centers {
		for _, vv := range v.Sessions {
			if (!options.Bookable || vv.AvailableCapacity > 0) && (options.Age == 0 || options.Age >= vv.MinAgeLimit) &&
				(options.Vaccine == "" || checkVaccine(options.Vaccine, vv.Vaccine)) {
				centerBookable = append(centerBookable, CenterBookable{
					Name:              v.Name,
					Freetype:          v.FeeType,
					SessionID:         vv.SessionID,
					Vaccine:           vv.Vaccine,
					MinAgeLimit:       vv.MinAgeLimit,
					Date:              vv.Date,
					AvailableCapacity: vv.AvailableCapacity,
				})

			}
		}
	}
	return centerBookable

}

// getSessionID gets session ID and generates OTP
func (scheduleData *ScheduleData) getSessionID(options Options, tokenValid bool) {

	var opt int
	centerBookable := getCenterBookable(options)

	if len(centerBookable) > 0 {
		// generate OTP only if there is bookable centers && invalid token
		if !tokenValid {
			scheduleData.txnId = genOTP(options.MobileNumber)
		}

		if options.Centers != "" {
			scheduleData.sessionID = getSpecifiedCenterSessionID(centerBookable, options.Centers)
		}

		if scheduleData.sessionID == "" {
			printCenterBookable(centerBookable)
			opt = getUserSelection("Enter Center ID :", len(centerBookable), false)

			scheduleData.sessionID = centerBookable[opt].SessionID
		}
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
		"captcha":       scheduleData.captcha,
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
	scheduleData.slot = options.Slot
	options.Bookable = true

	if runtime.GOOS == "android" && options.Aotp {
		_, lastRecievedTime = catchOTP()
	}

	if !options.Ntok {
		var ok bool
		scheduleData.bearerToken, ok = loadTokenFromFile()
		if ok {
			respCode, beneficaries = getBeneficaries(scheduleData.bearerToken)
			if respCode == 200 {
				tokenValid = true
			}
		}

	}

	scheduleData.getSessionID(options, tokenValid)

	if !tokenValid {
		if runtime.GOOS == "android" && options.Aotp {
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
			writeTokenToFile(scheduleData.bearerToken)
		}

		respCode, beneficaries = getBeneficaries(scheduleData.bearerToken)
		if respCode != 200 {
			log.Fatalln("Cannot get beneficaries")
		}
	}

	scheduleData.getBeneficariesID(beneficaries, options.Name)

	for i := 0; i < 5; i++ {

		scheduleData.captcha = getCatpchaCode(scheduleData.bearerToken)

		resp, statusCode := scheduleData.scheduleVaccineNow()

		switch statusCode {
		case 200:
			fmt.Println("Appointment scheduled successfully!")
			os.Exit(0)
		case 400:
			json.Unmarshal(resp, &badRequest)
			if badRequest.Error == "Your transaction didn't go through. Please try again later" {
				continue
			}
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

}
