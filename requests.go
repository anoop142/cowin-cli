package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func getReqAuth(URL string, bearerToken string) ([]byte, int) {
	client := http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	checkError(err)

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	resp, err := client.Do(req)

	checkError(err)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return body, resp.StatusCode

}

func postReq(URL string, postData []byte, bearerToken string) ([]byte, int) {
	client := http.Client{}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(postData))
	checkError(err)

	if bearerToken != "" {
		req.Header.Add("authorization", fmt.Sprintf("Bearer %s", bearerToken))
	}
	resp, err := client.Do(req)

	checkError(err)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body, resp.StatusCode
}