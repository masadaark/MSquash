package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:9999/healthcheck")
	if err != nil {
		//handle error
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) //ioutil => byte array
	fmt.Println(string(body))             //html pentex
}
