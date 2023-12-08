package credentials

import (
	"context"

	"github.com/strick-j/cybr-sdk-go/cybr"
)

const (
	// StaticCredentialName prvovides a name of Static Provider
	StaticCredentialsName = "StaticCredentials"
)

// StaticCredentialsEmptyError is emitted when static credentials are empty.
type StaticCredentialsEmptyError struct{}

func (*StaticCredentialsEmptyError) Error() string {
	return "static credentials are empty"
}

// A StaticCredentialsProvider is a set of credentials which are set, and will
// never expire.
type StaticCredentialsProvider struct {
	Value cybr.Credentials
}

// NewStaticCredentialsProvider return a StaticCredentialsProvider initialized with ISPSS
// credential passed in
func NewStaticCredentialsProvider(bearerToken string) StaticCredentialsProvider {
	return StaticCredentialsProvider{
		Value: cybr.Credentials{
			BearerToken: bearerToken,
		},
	}
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s StaticCredentialsProvider) Retrieve(_ context.Context) (cybr.Credentials, error) {
	v := s.Value
	if v.BearerToken == "" {
		return cybr.Credentials{
			Source: StaticCredentialsName,
		}, &StaticCredentialsEmptyError{}
	}

	if len(v.Source) == 0 {
		v.Source = StaticCredentialsName
	}

	return v, nil
}
