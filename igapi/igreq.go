package igapi

import (
	"io"
	"net/http"
)

var useragent = "Instagram 244.0.0.17.110 Android"

func GetRequest(requrl string) []byte {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", requrl, nil)
	// req.Header.Add("user-agent", useragent)
	// req.Header.Add("cookie", ``)

	resp, err := client.Do(req)
	chkerr(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	chkerr(err)

	return data
}
