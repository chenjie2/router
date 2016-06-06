// client project client.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", "http://127.0.0.1:8888/api", strings.NewReader("{\"command\":\"GetDomains\"}"))
	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}
