package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"main/endpoints"
	"main/tray"
	"net/http"
)

func main() {
	go handleRequests()
	systray.Run(tray.OnReady, tray.OnExit)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, endpoints.HandleEndPoints(r))
	fmt.Println(r.RequestURI)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":32000", nil))
}
