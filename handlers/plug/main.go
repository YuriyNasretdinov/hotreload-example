package main

import (
	"fmt"
	"net/http"
)

import "C"

func Example(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintf(rw, "Hello modified, %s!", req.Form.Get("name"))
}
