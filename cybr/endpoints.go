package cybr

import (
	"fmt"
)

// Endpoint represents the endpoint a service client should make API operation
// calls to.
//
// The SDK will automatically resolve these endpoints per API client using an
// internal endpoint resolvers. If you'd like to provide custom endpoint
// resolving behavior you can implement the EndpointResolver interface.
type Endpoint struct {
	// The base URL endpoint the SDK API clients will use to make API calls to.
	// The SDK will suffix URI path and query elements to this endpoint.
	URL string

	// Specifies if the endpoint's hostname can be modified by the SDK's API
	// client.
	//
	// If the hostname is mutable the SDK API clients may modify any part of
	// the hostname based on the requirements of the API, (e.g. adding, or
	// removing content in the hostname).
	//
	// Care should be taken when providing a custom endpoint for an API. If the
	// endpoint hostname is mutable, and the client cannot modify the endpoint
	// correctly, the operation call will most likely fail, or have undefined
	// behavior.
	//
	// If hostname is immutable, the SDK API clients will not modify the
	// hostname of the URL. This may cause the API client not to function
	// correctly if the API requires the operation specific hostname values
	// to be used by the client.
	//
	// This flag does not modify the API client's behavior if this endpoint
	// will be used instead of Endpoint Discovery, or if the endpoint will be
	// used to perform Endpoint Discovery. That behavior is configured via the
	// API Client's Options.
	HostnameImmutable bool

	// The AWS partition the endpoint belongs to.
	PartitionID string

	// The service name that should be used for signing the requests to the
	// endpoint.
	SigningName string

	// The signing method that should be used for signing the requests to the
	// endpoint.
	SigningMethod string

	// The source of the Endpoint. By default, this will be EndpointSourceServiceMetadata.
	// When providing a custom endpoint, you should set the source as EndpointSourceCustom.
	// If source is not provided when providing a custom endpoint, the SDK may not
	// perform required host mutations correctly. Source should be used along with
	// HostnameImmutable property as per the usage requirement.
	Source EndpointSource
}

// EndpointSource is the endpoint source type.
type EndpointSource int

// EndpointNotFoundError is a sentinel error to indicate that the
// EndpointResolver implementation was unable to resolve an endpoint for the
// given service and region. Resolvers should use this to indicate that an API
// client should fallback and attempt to use it's internal default resolver to
// resolve the endpoint.
type EndpointNotFoundError struct {
	Err error
}

// Error is the error message.
func (e *EndpointNotFoundError) Error() string {
	return fmt.Sprintf("endpoint not found, %v", e.Err)
}

// Unwrap returns the underlying error.
func (e *EndpointNotFoundError) Unwrap() error {
	return e.Err
}

// EndpointResolverWithOptions is an endpoint resolver that can be used to provide or
// override an endpoint for the given service, region, and the service client's EndpointOptions. API clients will
// attempt to use the EndpointResolverWithOptions first to resolve an endpoint if
// available. If the EndpointResolverWithOptions returns an EndpointNotFoundError error,
// API clients will fallback to attempting to resolve the endpoint using its
// internal default endpoint resolver.
type EndpointResolverWithOptions interface {
	ResolveEndpoint(service string, options ...interface{}) (Endpoint, error)
}

// EndpointResolverWithOptionsFunc wraps a function to satisfy the EndpointResolverWithOptions interface.
type EndpointResolverWithOptionsFunc func(service string, options ...interface{}) (Endpoint, error)

// ResolveEndpoint calls the wrapped function and returns the results.
func (e EndpointResolverWithOptionsFunc) ResolveEndpoint(service string, options ...interface{}) (Endpoint, error) {
	return e(service, options...)
}

// GetDisableHTTPS takes a service's EndpointResolverOptions and returns the DisableHTTPS value.
// Returns boolean false if the provided options does not have a method to retrieve the DisableHTTPS.
func GetDisableHTTPS(options ...interface{}) (value bool, found bool) {
	type iface interface {
		GetDisableHTTPS() bool
	}
	for _, option := range options {
		if i, ok := option.(iface); ok {
			value = i.GetDisableHTTPS()
			found = true
			break
		}
	}
	return value, found
}
