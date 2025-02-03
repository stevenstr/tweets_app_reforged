package discovery

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

var ErrNotFound = errors.New("no service addresses found")

// Registry defines the service registry.
type Registry interface {
	// Register creates a service instance record
	// in the registcry
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	// Deregister removes a service instance
	// from registry.
	Deregister(ctx context.Context, instanceID string, serviceName string) error
	// ServiceAddresses returns the list of addresses
	// of active instances of the given service.
	ServiceAddresses(ctx context.Context, serviceName string) ([]string, error)
	// ReportHealthyState is a push mechanism for reporting
	// healthy state to the registry.
	ReportHealthyState(instanceID string, _ string) error
}

// GenerateInstanceID generates a pseudo random
// service instance identifier using service name
// and rand number.
func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(uint64(time.Now().UnixNano()))).Int())
}
