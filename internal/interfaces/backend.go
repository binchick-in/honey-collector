package interfaces

import "context"

// HoneyBackend is the interface for publishing data to a backend.
type HoneyBackend interface {
	Publish(ctx context.Context, data []byte) error
}
