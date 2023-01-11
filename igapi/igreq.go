package igapi

import (
	"io"
	"net/http"
)

var useragent = "Mozilla/5.0 (Linux; Android 12; vivo 1939 Build/SP1A.210812.003; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/103.0.5060.71 Mobile Safari/537.36 Instagram 244.0.0.17.110 Android (31/12; 480dpi; 1080x2278; vivo; vivo 1939; 2004; qcom; en_US; 383877305)"

func GetRequest(requrl string) []byte {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", requrl, nil)
	req.Header.Add("user-agent", useragent)
	req.Header.Add("cookie", ``)

	resp, err := client.Do(req)
	chkerr(err)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	chkerr(err)

	return data
}
