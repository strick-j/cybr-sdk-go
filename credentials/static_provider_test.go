package credentials

import (
	"context"
	"testing"

	"github.com/strick-j/cybr-sdk-go/cybr"
)

func TestStaticCredentialsProvider(t *testing.T) {
	s := StaticCredentialsProvider{
		Value: cybr.Credentials{
			BearerToken: "BEARER",
		},
	}

	creds, err := s.Retrieve(context.Background())
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
	if e, a := "BEARER", creds.BearerToken; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestStaticCredentialsProviderIsExpired(t *testing.T) {
	s := StaticCredentialsProvider{
		Value: cybr.Credentials{
			BearerToken: "BEARER",
		},
	}

	creds, err := s.Retrieve(context.Background())
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if creds.Expired() {
		t.Errorf("expect static credentials to never expire")
	}
}
