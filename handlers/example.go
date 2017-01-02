package handlers

import (
	"fmt"
	"net/http"
	"os"
	"plugin"

	"github.com/bouk/monkey"
)

func Example(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintf(rw, "Hello, %s!", req.Form.Get("name"))
}

func Reload(rw http.ResponseWriter, req *http.Request) {
	p, err := plugin.Open("handlers.so")
	if err != nil {
		panic(err)
	}

	sym, err := p.Lookup("Example")
	if err != nil {
		panic(err)
	}

	os.Remove("handlers.so")

	monkey.Patch(Example, sym)

	fmt.Fprint(rw, "OK")
}
