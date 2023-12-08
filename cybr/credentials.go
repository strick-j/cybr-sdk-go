package cybr

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/strick-j/cybr-sdk-go/internal/sdk"
)

// AnonymousCredentials provides a sentinel CredentialsProvider that should be
// used to instruct the SDK's signing middleware to not sign the request.
//
// Using `nil` credentials when configuring an API client will achieve the same
// result. The AnonymousCredentials type allows you to configure the SDK's
// external config loading to not attempt to source credentials from the shared
// config or environment.
type AnonymousCredentials struct{}

// Retrieve implements the CredentialsProvider interface, but will always
// return error, and cannot be used to sign a request. The AnonymousCredentials
// type is used as a sentinel type instructing the AWS request signing
// middleware to not sign a request.
func (AnonymousCredentials) Retrieve(context.Context) (Credentials, error) {
	return Credentials{Source: "AnonymousCredentials"},
		fmt.Errorf("the AnonymousCredentials is not a valid credential provider, and cannot be used to sign AWS requests with")
}

type Credentials struct {
	// The Authorization Token.
	BearerToken string

	// Source of the credentials.
	Source string

	// States if the credentials can expire or not.
	CanExpire bool

	// The time when the credentials expire. Should be ignored when CanExpire is
	// false.
	Expires time.Time
}

// Expired returns if the credentials have expired.
func (v Credentials) Expired() bool {
	if v.CanExpire {
		// Calling Round(0) on the current time will truncate the monotonic
		// reading only. Ensures credential expiry time is always based on
		// reported wall-clock time.
		return !v.Expires.After(sdk.NowTime().Round(0))
	}

	return false
}

// HasKeys returns if the credentials keys are set.
func (v Credentials) HasKeys() bool {
	return len(v.BearerToken) > 0
}

// A CredentialsProvider is the interface for any component which will provide
// credentials Credentials. A CredentialsProvider is required to manage its own
// Expired state, and what to be expired means.
//
// A credentials provider implementation can be wrapped with a CredentialCache
// to cache the credential value retrieved. Without the cache the SDK will
// attempt to retrieve the credentials for every request.
type CredentialsProvider interface {
	// Retrieve returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve(ctx context.Context) (Credentials, error)
}

// CredentialsProviderFunc provides a helper wrapping a function value to
// satisfy the CredentialsProvider interface.
type CredentialsProviderFunc func(context.Context) (Credentials, error)

// Retrieve delegates to the function value the CredentialsProviderFunc wraps.
func (fn CredentialsProviderFunc) Retrieve(ctx context.Context) (Credentials, error) {
	return fn(ctx)
}

type isCredentialsProvider interface {
	IsCredentialsProvider(CredentialsProvider) bool
}

// IsCredentialsProvider returns whether the target CredentialProvider is the same type as provider when comparing the
// implementation type.
//
// If provider has a method IsCredentialsProvider(CredentialsProvider) bool it will be responsible for validating
// whether target matches the credential provider type.
//
// When comparing the CredentialProvider implementations provider and target for equality, the following rules are used:
//
//	If provider is of type T and target is of type V, true if type *T is the same as type *V, otherwise false
//	If provider is of type *T and target is of type V, true if type *T is the same as type *V, otherwise false
//	If provider is of type T and target is of type *V, true if type *T is the same as type *V, otherwise false
//	If provider is of type *T and target is of type *V,true if type *T is the same as type *V, otherwise false
func IsCredentialsProvider(provider, target CredentialsProvider) bool {
	if target == nil || provider == nil {
		return provider == target
	}

	if x, ok := provider.(isCredentialsProvider); ok {
		return x.IsCredentialsProvider(target)
	}

	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr {
		targetType = reflect.PtrTo(targetType)
	}

	providerType := reflect.TypeOf(provider)
	if providerType.Kind() != reflect.Ptr {
		providerType = reflect.PtrTo(providerType)
	}

	return targetType.AssignableTo(providerType)
}
