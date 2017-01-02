package main

import (
	"net/http"

	"github.com/YuriyNasretdinov/hotreload-example/handlers"
)

func main() {
	http.HandleFunc("/example", handlers.Example)
	http.HandleFunc("/reload", handlers.Reload)
	http.ListenAndServe(":9999", nil)
}
