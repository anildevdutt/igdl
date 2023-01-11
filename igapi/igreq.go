package igapi

import (
	"io"
	"net/http"
)

var useragent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

func GetRequest(requrl string) []byte {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", requrl, nil)
	req.Header.Add("user-agent", useragent)
	req.Header.Add("X-IG-App-ID", "936619743392459")
	req.Header.Add("cookie", ``)

	resp, err := client.Do(req)
	chkerr(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	chkerr(err)

	return data
}
