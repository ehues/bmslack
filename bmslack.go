package main

import (
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "9999"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	api := slack.New("xoxb-3965840334-lpTTUKO9Vm1QZLSBPKnGG0Fn")
	_, _, err := api.PostMessage("#general", "hello :globe_with_meridians:! 11:39.2", slack.PostMessageParameters{})
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	log.Printf("Successfully posted message")
}

func main() {
	var port string = os.Getenv("VCAP_APP_PORT")
	var host string = os.Getenv("VCAP_APP_HOST")

	if len(port) == 0 {
		port = DEFAULT_PORT
		log.Printf("No VCAP_APP_PORT. Using %+v\n", port)
	}

	if len(host) == 0 {
		host = DEFAULT_HOST
		log.Printf("No VCAP_APP_HOST. Using %+v\n", host)
	}

	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: ", err)
	}
}
