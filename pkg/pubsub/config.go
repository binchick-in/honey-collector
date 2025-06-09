package pubsub

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const GOOGLECLOUD_PLATFORM_CREDS = "GCP_CREDS"

var (
	ErrCredsNotFound = errors.New("could not find JSON creds in environment")
)

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

func getTopicFromEnv() string {
	topic, present := os.LookupEnv("PUBSUB_TOPIC")
	if !present {
		log.Println("PUBSUB_TOPIC environment variable not set, using default: honey")
		return "honey"
	}
	log.Printf("Using PUBSUB_TOPIC: %s", topic)
	return topic
}
