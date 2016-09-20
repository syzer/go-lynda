package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	res, err := http.Get("http://stackoverflow.com/questions/24780213/d3-js-how-to-rotate-the-texts-on-each-bar-of-a-grouped-bar-chart")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", page)
}
