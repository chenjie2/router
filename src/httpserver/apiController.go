package main

import (
	"fmt"
)

type apiController struct {
}

func (this *apiController) TestAction(paras map[string]string) string {
	return paras["command"]
}

func (this *apiController) Test1Action(paras map[string]string) string {
	return paras["command"]
}

func (this *apiController) GetDomainsAction(paras map[string]string) string {
	client := &kubeClient{"http://172.16.50.233"}
	result := client.GetDomains()
	return result
}

func (this *apiController) AddDomainAction(paras map[string]string) string {
	domaiName := paras["domain"]
	servers := paras["server"]
	fmt.Printf(domaiName + "=" + servers)
	client := &kubeClient{"http://172.16.50.233"}
	result := client.AddDomain(domaiName, servers)
	return result
}

func (this *apiController) DelDomainAction(paras map[string]string) string {
	domaiName := paras["domain"]
	client := &kubeClient{"http://172.16.50.233"}
	result := client.DelDomain(domaiName)
	return result
}
