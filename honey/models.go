package honey

import (
	"strings"
	"net/http"
	"encoding/json"
)

type LoggedRequest struct {
    Host string `json:"host"`
    Method string `json:"method"`
	Path string `json:"path"`
    RemoteAddress string `json:"remote_address"`
    UserAgent string `json:"user_agent"`
    QueryParams map[string]string `json:"query_params"`
    Headers map[string]string `json:"headers"`
    Body string `json:"body"`
}

func (lr *LoggedRequest) ToJson() string {
	j, err := json.Marshal(lr)
	if err != nil {
		panic(err)
	}
	return string(j)
}

func NewLoggedRequest(r http.Request) LoggedRequest {
	var decodedBody string
	if r.Method == http.MethodPost {
		decodedBody = decodeRequestBody(r.Body)
	}
	host := strings.Split(r.Host, ":")
	return LoggedRequest{
        Host:          host[0],
        Method:        r.Method,
		Path:          r.URL.Path,
        RemoteAddress: r.RemoteAddr,
        UserAgent:     r.UserAgent(),
        QueryParams:   processMapOfSlices(r.URL.Query()),
        Headers:       processMapOfSlices(r.Header),
        Body:          decodedBody,
    }

}
