package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	http.HandleFunc("/api", commandHandler)
	http.Handle("/", http.FileServer(http.Dir("webui")))
	http.ListenAndServe(":8888", nil)
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
		//		var f interface{}
		json.Unmarshal(result, &parameters)
		//		m := f.(map[string]string{})
		//		parameters = m
		//		for k, v := range m {
		//			switch vv := v.(type) {
		//			case string:
		//				fmt.Println(k, "is string", vv)
		//				parameters[k] = strings.Join(v, "")
		//			case int:
		//				fmt.Println(k, "is int", vv)
		//				parameters[k] = strings.Join(v, "")
		//			case float64:
		//				fmt.Println(k, "is float64", vv)
		//				parameters[k] = strings.Join(v, "")
		//			default:
		//				fmt.Println(k, "is of a type I don't know how to handle")
		//			}
		//		}
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
