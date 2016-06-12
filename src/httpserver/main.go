package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var port *string
var url *string

func main() {
	port = flag.String("port", "8888", "http server port")
	url = flag.String("url", "https://domain.yysoma.com", "router url")
	http.HandleFunc("/api", commandHandler)
	http.Handle("/", http.FileServer(http.Dir("webui")))
	http.ListenAndServe(":"+*port, nil)
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	parameters := make(map[string]string)
	controller := &apiController{}
	if r.Method == "GET" {
		for k, v := range r.Form {
			parameters[k] = strings.Join(v, "")
			fmt.Print("key:", k, "; ")
			fmt.Println("val:", strings.Join(v, ""))
		}
	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		json.Unmarshal(result, &parameters)
	}
	commandName := parameters["command"]
	if commandName == "" {
		w.Write([]byte("command not found"))
		return
	}
	fmt.Println(commandName)
	commandAction := commandName + "Action"
	admin := reflect.ValueOf(controller)
	action := admin.MethodByName(commandAction)
	para := reflect.ValueOf(parameters)
	result := action.Call([]reflect.Value{para})
	w.Write([]byte(result[0].String()))
}
