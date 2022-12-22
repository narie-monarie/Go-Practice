package main

import (
	"fmt"
	"net/http"
	"time"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The current time is %s", time.Now())
}

func requestSay(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("name")

	if val != "" {
		fmt.Fprintf(w, "Hello %s", val)
	} else {
		fmt.Fprintf(w, "Hello ... you")
	}
}
func main() {
	fmt.Println("starting on port :3000")
	http.HandleFunc("/time", requestHandler)
	http.HandleFunc("/say", requestSay)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
