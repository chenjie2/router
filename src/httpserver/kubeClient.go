package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type kubeClient struct {
	Url string
}

type Upstream struct {
	DomainName string
	Server     string
}

func (this *kubeClient) GetNode() string {
	json := `{"key1":"value1","key2":"value2","key3":"value3"}`
	return json
}

func (thbis *kubeClient) AddDomain(domainName string, servers string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, _ := http.NewRequest("POST", "https://domain.yysoma.com/upstream/"+domainName, strings.NewReader(servers))
	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		return bodystr
	}
	return "failed"
}

func (this *kubeClient) DelDomain(domainName string) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, _ := http.NewRequest("DELETE", "https://domain.yysoma.com/upstream/"+domainName, nil)
	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
		return bodystr
	}
	return "failed"
}

func (this *kubeClient) GetDomains() string {
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
	}
	jsonstr, _ := json.Marshal(silce)
	fmt.Println(string(jsonstr))
	return string(jsonstr)
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	return string(body)

}
