package honey

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var ctx = context.Background()

type hostMetaData map[string]string

type HoneyClient struct {
	client       *pubsub.Client
	topic        *pubsub.Topic
	hostMetaData hostMetaData
}

func (h *HoneyClient) Publish(data []byte) {
	msg := pubsub.Message{
		Data:       data,
		Attributes: h.hostMetaData,
	}
	h.topic.Publish(ctx, &msg)
}

func NewHoneyClientFromEnv() HoneyClient {
	creds, projectid := checkEnvForCreds()
	pubSubClient, pubSubClientErr := pubsub.NewClient(
		ctx,
		projectid,
		option.WithCredentialsJSON(creds),
	)
	if pubSubClientErr != nil {
		panic(pubSubClientErr)
	}

	hostname, hostnameErr := os.Hostname()
	if hostnameErr != nil {
		panic(hostnameErr)
	}

	return HoneyClient{
		client: pubSubClient,
		topic:  pubSubClient.Topic("honey"), // TODO: Put topic string in environment
		hostMetaData: hostMetaData{
			"hostname": hostname,
		},
	}
}
