package cowin

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getReqAuth(URL, bearerToken string, auth bool) ([]byte, int) {
	var body []byte
	var respCode int
	client := http.Client{}

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Fatalln(err.Error())
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux x86_64) Chrome/90.0.4430.93 Safari/537.36")
	if auth {
		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	}
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		body, _ = io.ReadAll(resp.Body)
		respCode = resp.StatusCode
	}

	return body, respCode

}

func postReq(URL string, postData []byte, bearerToken string) ([]byte, int) {
	var body []byte
	var respCode int

	client := http.Client{}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(postData))
	if err != nil {
		log.Fatalln(err.Error())
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Linux x86_64) Chrome/90.0.4430.93 Safari/537.36")

	if bearerToken != "" {
		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	}
	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
		body, _ = io.ReadAll(resp.Body)
		respCode = resp.StatusCode
	}

	return body, respCode
}
