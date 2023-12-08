package cybr

import (
	"net/http"

	smithybearer "github.com/aws/smithy-go/auth/bearer"
	"github.com/aws/smithy-go/logging"
	"github.com/aws/smithy-go/middleware"
)

// HTTPClient provides the interface to provide custom HTTPClients. Generally
// *http.Client is sufficient for most use cases. The HTTPClient should not
// follow 301 or 302 redirects.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// A Config provides service configuration for service clients.
type Config struct {
	// The tenant name to use for the service client. This parameter is required
	// and must configured globally or on a per-client basis unless otherwise noted.
	TenantName string

	// The tenant ID to use for the service client. This parameter is optional
	// and may be configured globally or on a per-client basis. If not provided,
	// the tenant ID will be resolved using the tenant name.
	TenantID string

	// The credentials object to use when signing requests.
	// Use the LoadDefaultConfig to load configuration from all the SDK's supported
	// sources, and resolve credentials using the SDK's default credential chain.
	Credentials CredentialsProvider

	// The Bearer Authentication token provider to use for authenticating API
	// operation calls with a Bearer Authentication token. The API clients and
	// operation must support Bearer Authentication scheme in order for the
	// token provider to be used. API clients created with NewFromConfig will
	// automatically be configured with this option, if the API client support
	// Bearer Authentication.
	//
	// The SDK's config.LoadDefaultConfig can automatically populate this
	// option for external configuration options such as SSO session.
	BearerAuthTokenProvider smithybearer.TokenProvider

	// The HTTP Client the SDK's API clients will use to invoke HTTP requests.
	// The SDK defaults to a BuildableClient allowing API clients to create
	// copies of the HTTP Client for service specific customizations.
	//
	// Use a (*http.Client) for custom behavior. Using a custom http.Client
	// will prevent the SDK from modifying the HTTP client.
	HTTPClient HTTPClient

	// RetryMaxAttempts specifies the maximum number attempts an API client
	// will call an operation that fails with a retryable error.
	//
	// API Clients will only use this value to construct a retryer if the
	// Config.Retryer member is not nil. This value will be ignored if
	// Retryer is not nil.
	RetryMaxAttempts int

	// RetryMode specifies the retry model the API client will be created with.
	//
	// API Clients will only use this value to construct a retryer if the
	// Config.Retryer member is not nil. This value will be ignored if
	// Retryer is not nil.
	RetryMode RetryMode

	// Retryer is a function that provides a Retryer implementation. A Retryer
	// guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	//
	// In general, the provider function should return a new instance of a
	// Retryer if you are attempting to provide a consistent Retryer
	// configuration across all clients. This will ensure that each client will
	// be provided a new instance of the Retryer implementation, and will avoid
	// issues such as sharing the same retry token bucket across services.
	//
	// If not nil, RetryMaxAttempts, and RetryMode will be ignored by API
	// clients.
	Retryer func() Retryer

	// ConfigSources are the sources that were used to construct the Config.
	// Allows for additional configuration to be loaded by clients.
	ConfigSources []interface{}

	// APIOptions provides the set of middleware mutations modify how the API
	// client requests will be handled. This is useful for adding additional
	// tracing data to a request, or changing behavior of the SDK's client.
	APIOptions []func(*middleware.Stack) error

	// The logger writer interface to write logging messages to. Defaults to
	// standard error.
	Logger logging.Logger

	// AppId is an optional application specific identifier that can be set.
	// When set it will be appended to the User-Agent header of every request
	// in the form of App/{AppId}. This variable is sourced from environment
	// variable AWS_SDK_UA_APP_ID or the shared config profile attribute sdk_ua_app_id.
	// TODO: EDIT ABOVE
	AppID string

	// BaseEndpoint is an intermediary transfer location to a service specific
	// BaseEndpoint on a service's Options.
	BaseEndpoint *string
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewConfig() *Config {
	return &Config{}
}

// Copy will return a shallow copy of the Config object. If any additional
// configurations are provided they will be merged into the new config returned.
func (c Config) Copy() Config {
	cp := c
	return cp
}

// EndpointDiscoveryEnableState indicates if endpoint discovery is
// enabled, disabled, auto or unset state.
//
// Default behavior (Auto or Unset) indicates operations that require endpoint
// discovery will use Endpoint Discovery by default. Operations that
// optionally use Endpoint Discovery will not use Endpoint Discovery
// unless EndpointDiscovery is explicitly enabled.
type EndpointDiscoveryEnableState uint

// Enumeration values for EndpointDiscoveryEnableState
const (
	// EndpointDiscoveryUnset represents EndpointDiscoveryEnableState is unset.
	// Users do not need to use this value explicitly. The behavior for unset
	// is the same as for EndpointDiscoveryAuto.
	EndpointDiscoveryUnset EndpointDiscoveryEnableState = iota

	// EndpointDiscoveryAuto represents an AUTO state that allows endpoint
	// discovery only when required by the api. This is the default
	// configuration resolved by the client if endpoint discovery is neither
	// enabled or disabled.
	EndpointDiscoveryAuto // default state

	// EndpointDiscoveryDisabled indicates client MUST not perform endpoint
	// discovery even when required.
	EndpointDiscoveryDisabled

	// EndpointDiscoveryEnabled indicates client MUST always perform endpoint
	// discovery if supported for the operation.
	EndpointDiscoveryEnabled
)
