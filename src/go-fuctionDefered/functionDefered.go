package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// makes request to get robots.txt file from google.com
	// file is response
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Close resource last
	defer res.Body.Close()
	// ReadAll func  = take stream res.Body and parse out bites
	robots, err := ioutil.ReadAll(res.Body)
	 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}