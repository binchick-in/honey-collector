package pubsub

import (
	"context"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"

	"honey-collector/internal/interfaces"
)

type hostMetaData map[string]string

type GooglePubSubClient struct {
	client       *pubsub.Client
	topic        *pubsub.Topic
	hostMetaData hostMetaData
}

func (h *GooglePubSubClient) Publish(ctx context.Context, data []byte) error {
	msg := &pubsub.Message{
		Data:       data,
		Attributes: h.hostMetaData,
	}
	result := h.topic.Publish(ctx, msg)
	_, err := result.Get(ctx)
	return err
}

func NewGooglePubSubClient() (interfaces.HoneyBackend, error) {
	creds, projectid, err := checkEnvForCreds()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	pubSubClient, pubSubClientErr := pubsub.NewClient(ctx, projectid, option.WithCredentialsJSON(creds))
	if pubSubClientErr != nil {
		return nil, pubSubClientErr
	}

	hostname, hostnameErr := os.Hostname()
	if hostnameErr != nil {
		return nil, hostnameErr
	}

	return &GooglePubSubClient{
		client: pubSubClient,
		topic:  pubSubClient.Topic(getTopicFromEnv()),
		hostMetaData: hostMetaData{
			"hostname": hostname,
		},
	}, nil
}