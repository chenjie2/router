package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

//type Upstream struct {
//	DomainName string
//	Server     string
//}

func Test_httpGet(t *testing.T) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://domain.yysoma.com/detail")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	details := strings.Split(string(body), "\n")
	silce := []Upstream{}
	var domainName string
	var servers string
	for _, value := range details {
		if !strings.HasPrefix(value, "server") && value != "" {
			domainName = value
		} else if strings.HasPrefix(value, "server") && value != "" {
			servers += value + ";"
		} else if value == "" {
			if domainName != "" {
				silce = append(silce, Upstream{DomainName: domainName, Server: servers})
			}
			domainName = ""
			servers = ""
		}
		//		domain := &upstream{}
		//		domain.domainName = value
		//		fmt.Println(value)
		//		resp, _ := client.Get("https://domain.yysoma.com/upstream/" + value)
		//		upstream, _ := ioutil.ReadAll(resp.Body)
		//		fmt.Println(string(upstream))
		//		domain.server = strings.Replace(string(upstream), "\n", ";", -1)
		//		fmt.Println(domain)
		//		jsonstr, _ := json.Marshal(domain)
		//		fmt.Printf("Value: %v\n", string(jsonstr))
		//		silce = append(silce, Upstream{DomainName: value, Server: strings.Replace(string(upstream), "\n", ";", -1)})
	}
	jsonstr, _ := json.Marshal(silce)
	fmt.Println(string(jsonstr))

}
