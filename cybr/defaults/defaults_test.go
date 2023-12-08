package defaults

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/strick-j/cybr-sdk-go/cybr"

	"github.com/google/go-cmp/cmp"
)

// v1TestResolver returns the default Configuration descriptor for the given mode.
//
// Supports the following modes: cross-region, in-region, mobile, standard
func v1TestResolver(mode cybr.DefaultsMode) (Configuration, error) {
	var mv cybr.DefaultsMode
	mv.SetFromString(string(mode))

	switch mv {
	case cybr.DefaultsModeMobile:
		settings := Configuration{
			ConnectTimeout:        cybr.Duration(10000 * time.Millisecond),
			RetryMode:             cybr.RetryMode("adaptive"),
			TLSNegotiationTimeout: cybr.Duration(11000 * time.Millisecond),
		}
		return settings, nil
	case cybr.DefaultsModeStandard:
		settings := Configuration{
			ConnectTimeout:        cybr.Duration(2000 * time.Millisecond),
			RetryMode:             cybr.RetryMode("standard"),
			TLSNegotiationTimeout: cybr.Duration(2000 * time.Millisecond),
		}
		return settings, nil
	default:
		return Configuration{}, fmt.Errorf("unsupported defaults mode: %v", mode)
	}
}

func TestConfigV1(t *testing.T) {
	cases := []struct {
		Mode     cybr.DefaultsMode
		Expected Configuration
	}{
		{
			Mode: cybr.DefaultsModeStandard,
			Expected: Configuration{
				ConnectTimeout:        cybr.Duration(2000 * time.Millisecond),
				TLSNegotiationTimeout: cybr.Duration(2000 * time.Millisecond),
				RetryMode:             cybr.RetryModeStandard,
			},
		},
		{
			Mode: cybr.DefaultsModeMobile,
			Expected: Configuration{
				ConnectTimeout:        cybr.Duration(10000 * time.Millisecond),
				TLSNegotiationTimeout: cybr.Duration(11000 * time.Millisecond),
				RetryMode:             cybr.RetryModeAdaptive,
			},
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := v1TestResolver(tt.Mode)
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
			if diff := cmp.Diff(tt.Expected, got); len(diff) > 0 {
				t.Error(diff)
			}
		})
	}
}
