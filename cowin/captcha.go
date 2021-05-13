package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

const (
	captchaImageFile    = "captcha.svg"
	captchaImageFilePng = "captcha.png"
	imgViewer           = "pixterm"
	svg2pngConverter    = "convert"
)

func cleanCaptchaImg(img string) []byte {
	reg := regexp.MustCompile(`(<path d=)(.*?)(fill=\"none\"/>)`)
	return []byte(reg.ReplaceAllString(img, ""))
}

func writeCaptchaImg(bearerToken string) bool {
	emptyData := map[string]string{}
	postData, _ := json.Marshal(emptyData)

	resp, responseCode := postReq(captchaURL, postData, bearerToken)

	if responseCode != 200 {
		log.Fatalln("Cannot get captcha")
	}
	var captchaData map[string]string

	json.Unmarshal(resp, &captchaData)

	captchaImg, ok := captchaData["captcha"]

	if !ok {
		log.Fatalln("Cannot get captcha Image")
	}

	err := os.WriteFile(captchaImageFile, cleanCaptchaImg(captchaImg), 0644)

	return err == nil

}

// check for programs installed to render captcha in terminal
func checkImageTerminalDep() bool {
	dep := []string{svg2pngConverter, imgViewer}
	stsf := 0

	for _, v := range dep {
		_, err := exec.LookPath(v)

		if err != nil {
			break
		}
		stsf++
	}
	return stsf == len(dep)

}

// Linux Only!
func displayCaptchaImageTerminal() {
	imgViewerParams := []string{"-tr", "16", captchaImageFilePng}
	//convert svg to png
	cmd := exec.Command(svg2pngConverter, captchaImageFile, captchaImageFilePng)
	cmd.Run()
	// view the captcha image
	cmd = exec.Command(imgViewer, imgViewerParams...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Windows WIP
func displayCaptchaImageWeb() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C "+captchaImageFile)
	} else {
		cmd = exec.Command("firefox", captchaImageFile)
	}
	cmd.Start()

}

// platform dependent
func displayCaptchaImage() {

	if checkImageTerminalDep() {
		displayCaptchaImageTerminal()
	} else if runtime.GOOS == "android" {
		fmt.Println("Captcha rendering dependencies missing")
		os.Exit(1)
	} else {
		displayCaptchaImageWeb()
	}
}

func userInputCaptcha() string {
	// initialise to null
	captcha := "null"
	fmt.Printf("Enter captcha: ")
	fmt.Scanf("%s\n", &captcha)

	return captcha
}
