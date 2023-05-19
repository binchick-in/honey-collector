package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"

	"honey-collector/honey"
)

var (
	honeyClient  honey.HoneyClient
	ports        string
	responseText string
)

func ReqHandler(resp http.ResponseWriter, req *http.Request) {
	lr := honey.NewLoggedRequest(*req)
	fmt.Println(lr.ToJson())
	honeyClient.Publish([]byte(lr.ToJson()))
	fmt.Fprint(resp, responseText)
}

func main() {
	flag.StringVar(&ports, "ports", "", "Wrap your port list in double quotes")
	flag.StringVar(&responseText, "response", "\\( ^ o ^)/", "String to respond to web requests with")
	flag.Parse()

	var wg sync.WaitGroup
	preppedPorts := honey.PreparePorts(ports)
	honeyClient = honey.NewHoneyClientFromEnv()
	http.HandleFunc("/", ReqHandler)

	for _, p := range preppedPorts {
		wg.Add(1)
		addr := fmt.Sprintf(":%s", p)
		fmt.Printf("Starting honey pot on port: %s\n", p)
		go http.ListenAndServe(addr, nil)
	}
	wg.Wait()
}
