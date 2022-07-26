package main

import (
	"fmt"
	"net/http"

	"honey-collector/honey"
)

var honeyClient honey.HoneyClient

func ReqHandler(resp http.ResponseWriter, req *http.Request) {
	lr := honey.NewLoggedRequest(*req)
	fmt.Println(lr.ToJson())
	fmt.Fprintf(resp, "ree")
}

func main() {
	defaultPort := 8081
	honeyClient = honey.NewHoneyClientFromEnv()
	fmt.Printf("Starting honey pot on port: %d\n", defaultPort)
	http.HandleFunc("/", ReqHandler)
	serverStartErr := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)
	if serverStartErr != nil {
		panic(serverStartErr)
	}
}
