package honey

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const GOOGLECLOUD_PLATFORM_CREDS = "GCP_CREDS"

var (
	ErrCredsNotFound = errors.New("could not find JSON creds in environment")
)

/*
Check the environment for GCP_CREDS and then parse and return
*/
func checkEnvForCreds() (creds []byte, projectId string) {
	var structuredCreds map[string]string
	val, present := os.LookupEnv(GOOGLECLOUD_PLATFORM_CREDS)
	if !present {
		panic(ErrCredsNotFound)
	}
	creds = []byte(val)
	jErr := json.Unmarshal(creds, &structuredCreds)
	if jErr != nil {
		panic(jErr)
	}
	projectId = structuredCreds["project_id"] // Handle missing key
	return
}

func processMapOfSlices(x map[string][]string) map[string]string {
	r := make(map[string]string)
	for k, v := range x {
		r[k] = v[0]
	}
	return r
}

func decodeRequestBody(b io.ReadCloser) string {
	rawBody, err := ioutil.ReadAll(b)
	if err != nil {
		log.Print(err)
		return "ERROR PARSING BODY"
	}
	return string(rawBody)
}

func PreparePorts(x string) (a []string) {
	ports := strings.Split(x, ",")
	for _, i := range ports {
		a = append(a, strings.TrimSpace(i))
	}
	return
}
