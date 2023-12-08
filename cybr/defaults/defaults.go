package defaults

import (
	"fmt"
	"time"

	"github.com/strick-j/cybr-sdk-go/cybr"
)

// GetModeConfiguration returns the default Configuration descriptor for the given mode.
//
// Supports the following modes: cross-region, in-region, mobile, standard
func GetModeConfiguration(mode cybr.DefaultsMode) (Configuration, error) {
	var mv cybr.DefaultsMode
	mv.SetFromString(string(mode))

	switch mv {
	case cybr.DefaultsModeMobile:
		settings := Configuration{
			ConnectTimeout:        cybr.Duration(30000 * time.Millisecond),
			RetryMode:             cybr.RetryMode("standard"),
			TLSNegotiationTimeout: cybr.Duration(30000 * time.Millisecond),
		}
		return settings, nil
	case cybr.DefaultsModeStandard:
		settings := Configuration{
			ConnectTimeout:        cybr.Duration(3100 * time.Millisecond),
			RetryMode:             cybr.RetryMode("standard"),
			TLSNegotiationTimeout: cybr.Duration(3100 * time.Millisecond),
		}
		return settings, nil
	default:
		return Configuration{}, fmt.Errorf("unsupported defaults mode: %v", mode)
	}
}
