package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

const GPIO string = "" //GPIO pin number (string type)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.html")
	CheckErr(err)
	if r.Method == "GET" {
		stat, err := ioutil.ReadFile("/sys/class/gpio/gpio" + GPIO + "/value")
		CheckErr(err)
		if string(stat[0]) == "1" {
			temp.Execute(w, "checked")
			return
		} else if string(stat[0]) == "0" {
			temp.Execute(w, nil)
			return
		}
	}

	r.ParseForm()
	if r.PostForm.Get("power") != "" {
		CheckErr(ioutil.WriteFile("/sys/class/gpio/gpio17/value", []byte("1"), 0666))
		temp.Execute(w, "checked")
		return
	}
	CheckErr(ioutil.WriteFile("/sys/class/gpio/gpio17/value", []byte("0"), 0666))
	temp.Execute(w, nil)
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
