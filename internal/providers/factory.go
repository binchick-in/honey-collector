package providers

import (
	"errors"

	"honey-collector/internal/interfaces"
	"honey-collector/pkg/pubsub"
)

func NewHoneyBackend(backendName string) (interfaces.HoneyBackend, error) {
	switch backendName {
	case "google":
		return pubsub.NewGooglePubSubClient()
	case "sql":
		// TODO: Implement SQL client
		return nil, errors.New("SQL client not implemented yet")
	default:
		return nil, errors.New("unsupported provider: " + backendName)
	}
}
