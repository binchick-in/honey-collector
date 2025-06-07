package honey

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

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
