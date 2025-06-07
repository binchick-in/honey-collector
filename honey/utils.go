package honey

import (
	"encoding/json"
	"errors"
	"io"
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
func checkEnvForCreds() ([]byte, string, error) {
	var structuredCreds map[string]string
	val, present := os.LookupEnv(GOOGLECLOUD_PLATFORM_CREDS)
	if !present {
		return nil, "", ErrCredsNotFound
	}
	creds := []byte(val)
	if err := json.Unmarshal(creds, &structuredCreds); err != nil {
		return nil, "", err
	}
	projectId, ok := structuredCreds["project_id"]
	if !ok {
		return nil, "", errors.New("project_id not found in credentials")
	}
	return creds, projectId, nil
}

func processMapOfSlices(x map[string][]string) map[string]string {
	r := make(map[string]string)
	for k, v := range x {
		r[k] = v[0]
	}
	return r
}

func decodeRequestBody(b io.ReadCloser) string {
	rawBody, err := io.ReadAll(b)
	if err != nil {
		log.Print(err)
		return "ERROR PARSING BODY"
	}
	return string(rawBody)
}

func getTopicFromEnv() string {
	topic, present := os.LookupEnv("PUBSUB_TOPIC")
	if !present {
		return "default-topic" // Default topic if not set
	}
	return topic
}

func PreparePorts(x string) (a []string) {
	ports := strings.Split(x, ",")
	for _, i := range ports {
		a = append(a, strings.TrimSpace(i))
	}
	return
}
