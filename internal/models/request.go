package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// LoggedRequest represents a structured HTTP request for logging
type LoggedRequest struct {
	Host          string            `json:"host"`
	Method        string            `json:"method"`
	Path          string            `json:"path"`
	RemoteAddress string            `json:"remote_address"`
	UserAgent     string            `json:"user_agent"`
	QueryParams   map[string]string `json:"query_params"`
	Headers       map[string]string `json:"headers"`
	Body          string            `json:"body"`
	Time          uint64            `json:"time"`
}

func (lr *LoggedRequest) ToJson() (string, error) {
	j, err := json.Marshal(lr)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func NewLoggedRequest(r *http.Request) LoggedRequest {
	var decodedBody string
	if r.Method == http.MethodPost {
		decodedBody = decodeRequestBody(r.Body)
	}
	remoteHost := strings.SplitN(r.RemoteAddr, ":", 2)
	return LoggedRequest{
		Host:          r.Host,
		Method:        r.Method,
		Path:          r.URL.Path,
		RemoteAddress: remoteHost[0],
		UserAgent:     r.UserAgent(),
		QueryParams:   processMapOfSlices(r.URL.Query()),
		Headers:       processMapOfSlices(r.Header),
		Body:          decodedBody,
		Time:          uint64(time.Now().Unix()),
	}
}

// Converts a map with string keys and slice of strings values to a map with string keys and string values.
func processMapOfSlices(x map[string][]string) map[string]string {
	r := make(map[string]string)
	for k, v := range x {
		r[k] = v[0]
	}
	return r
}

// NOTE: There seems to be many instances of errors here, but I'm unsure what kind of body is causing an error on read
func decodeRequestBody(b io.ReadCloser) string {
	rawBody, err := io.ReadAll(b)
	if err != nil {
		log.Print(err)
		return "ERROR PARSING BODY"
	}
	return string(rawBody)
}
