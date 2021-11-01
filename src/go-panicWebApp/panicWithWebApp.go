package main

import (
	"net/http"
)

func main() {
	//Web App
	//HandleFunc listener - will listen on every URL (/).
	// func is a callback that gets called everytime a route comes in -
	// when this happens the string "Hello Don" is written to the browser
	//
	// TO TEST : launch go program and bring up BROWSER and enter URL
	// localhost:8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Don"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
