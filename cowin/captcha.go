package cowin

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const (
	captchaImageFile    = "captcha.svg"
	captchaImageFilePng = "captcha.png"
	imgViewer           = "pixterm"
)

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

	err := os.WriteFile(captchaImageFile, []byte(captchaImg), 0644)

	return err == nil

}

// Linux Only!
func displayCaptchaImageTerminal() {
	imgViewerParams := []string{"-tr", "16", captchaImageFilePng}
	//convert svg to png
	cmd := exec.Command("convert", captchaImageFile, captchaImageFilePng)
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

func userInputCaptcha() string {
	// initialise to null
	captcha := "null"
	fmt.Printf("Enter captcha: ")
	fmt.Scanf("%s\n", &captcha)

	return captcha
}
