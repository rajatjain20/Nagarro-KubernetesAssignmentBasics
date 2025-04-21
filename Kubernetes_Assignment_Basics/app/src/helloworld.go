package main

import (
	"fmt"
	"net/http"
	"time"
)

var started time.Time

func main() {
	fmt.Println("Hello World application")

	started = time.Now()

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/healthz", getHealth)

	fmt.Println("Listning at port :3000")
	http.ListenAndServe(":3000", nil)
}
