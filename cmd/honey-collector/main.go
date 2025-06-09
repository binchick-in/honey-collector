// Main entry point for the honey-collector application.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"honey-collector/internal/interfaces"
	"honey-collector/internal/models"
	"honey-collector/internal/providers"
	"honey-collector/pkg/utils"
)

var (
	ports        string
	responseText string
	backendName string
)

func ReqHandler(honeyBackend interfaces.HoneyBackend, resp http.ResponseWriter, req *http.Request) {
	lr := models.NewLoggedRequest(req)
	jsonStr, err := lr.ToJson()
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Request: %s", jsonStr)
	if err := honeyBackend.Publish(req.Context(), []byte(jsonStr)); err != nil {
		log.Printf("Publish error: %v", err)
		http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(resp, responseText)
}

// Starr the HTTP server listening on the specified port
func startListener(honeyClient interfaces.HoneyBackend, startErrorChannel chan<- error, port string) {
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting honey pot on port: %s", port)
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		ReqHandler(honeyClient, resp, req)
	})
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		startErrorChannel <- err
	}
}

func main() {
	flag.StringVar(&responseText, "response", "\\( ^ o ^)/", "String to respond to web requests with")
	flag.StringVar(&ports, "ports", "", "Comma-separated list of ports")
	flag.StringVar(&backendName, "provider", "sql", "Name of the data processing provider: sql, google, etc")
	flag.Parse()

	preppedPorts := utils.PreparePorts(ports)
	honeyBackend, err := providers.NewHoneyBackend(backendName)
	if err != nil {
		log.Fatalf("Failed to create Honey Provider: %v", err)
	}
	startErrorChannel := make(chan error, len(preppedPorts))

	for _, p := range preppedPorts {
		go startListener(honeyBackend, startErrorChannel, p)
	}
	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	select {
	case err := <-startErrorChannel:
		log.Fatalf("Server error: %v", err)
	case sig := <-sigChan:
		log.Printf("Received signal: %v. Shutting down.", sig)
	}
}