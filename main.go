package main

import (
	"log"
	"flag"
	"fmt"
	"net/http"

	"honey-collector/honey"
)

var (
	honeyClient  honey.HoneyClient
	ports        string
	responseText string
)

func ReqHandler(resp http.ResponseWriter, req *http.Request) {
	lr := honey.NewLoggedRequest(*req)
	log.Println(lr.ToJson())
	honeyClient.Publish([]byte(lr.ToJson()))
	fmt.Fprint(resp, responseText)
}


func startListener(startErrorChannel chan<- error, port string) {
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Starting honey pot on port: %s\n", port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		startErrorChannel <- err
	}
}

func main() {
	flag.StringVar(&ports, "ports", "", "Wrap your port list in double quotes")
	flag.StringVar(&responseText, "response", "\\( ^ o ^)/", "String to respond to web requests with")
	flag.Parse()

	preppedPorts := honey.PreparePorts(ports)
	honeyClient = honey.NewHoneyClientFromEnv()
	http.HandleFunc("/", ReqHandler)
	startErrorChannel := make(chan error)

	for _, p := range preppedPorts {
		go startListener(startErrorChannel, p)
	}
	panic(<-startErrorChannel)
}
