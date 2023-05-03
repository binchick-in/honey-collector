package honey

import (
	"context"
	"fmt"
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
	r := h.topic.Publish(ctx, &msg) // Leave this unassigned for now. We might want to handle this later though.
	s, err := r.Get(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(s) // Remove this line
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
