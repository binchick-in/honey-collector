package main

import (
    "fmt"
    "net/http"

    "honey-collector/honey"
)

var honeyClient honey.HoneyClient

func ReqHandler(resp http.ResponseWriter, req *http.Request) {
    lr := honey.NewLoggedRequest(*req)
    honeyClient.Publish([]byte(lr.ToJson()))
    fmt.Fprintf(resp, "\\( ^ o ^)/")
}

func main() {
    defaultPort := 80
    honeyClient = honey.NewHoneyClientFromEnv()
    fmt.Printf("Starting honey pot on port: %d\n", defaultPort)
    http.HandleFunc("/", ReqHandler)
    serverStartErr := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)
    if serverStartErr != nil {
        panic(serverStartErr)
    }
}
