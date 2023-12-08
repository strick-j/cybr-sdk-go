package cybr

import (
	"strings"
)

// DefaultsMode is the SDK defaults mode setting.
type DefaultsMode string

// The DefaultsMode constants.
const (
	// DefaultsModeAuto is an experimental mode that builds on the standard mode.
	// The SDK will attempt to discover the execution environment to determine the
	// appropriate settings automatically.
	//
	// Note that the auto detection is heuristics-based and does not guarantee 100%
	// accuracy. STANDARD mode will be used if the execution environment cannot
	// be determined. The auto detection might query EC2 Instance Metadata service
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html),
	// which might introduce latency. Therefore we recommend choosing an explicit
	// defaults_mode instead if startup latency is critical to your application
	DefaultsModeAuto DefaultsMode = "auto"

	// DefaultsModeLegacy provides default settings that vary per SDK and were used
	// prior to establishment of defaults_mode
	DefaultsModeLegacy DefaultsMode = "legacy"

	// DefaultsModeMobile builds on the standard mode and includes optimization
	// tailored for mobile applications
	//
	// Note that the default values vended from this mode might change as best practices
	// may evolve. As a result, it is encouraged to perform tests when upgrading
	// the SDK
	DefaultsModeMobile DefaultsMode = "mobile"

	// DefaultsModeStandard provides the latest recommended default values that
	// should be safe to run in most scenarios
	//
	// Note that the default values vended from this mode might change as best practices
	// may evolve. As a result, it is encouraged to perform tests when upgrading
	// the SDK
	DefaultsModeStandard DefaultsMode = "standard"
)

// SetFromString sets the DefaultsMode value to one of the pre-defined constants that matches
// the provided string when compared using EqualFold. If the value does not match a known
// constant it will be set to as-is and the function will return false. As a special case, if the
// provided value is a zero-length string, the mode will be set to LegacyDefaultsMode.
func (d *DefaultsMode) SetFromString(v string) (ok bool) {
	switch {
	case strings.EqualFold(v, string(DefaultsModeAuto)):
		*d = DefaultsModeAuto
		ok = true
	case strings.EqualFold(v, string(DefaultsModeLegacy)):
		*d = DefaultsModeLegacy
		ok = true
	case strings.EqualFold(v, string(DefaultsModeMobile)):
		*d = DefaultsModeMobile
		ok = true
	case strings.EqualFold(v, string(DefaultsModeStandard)):
		*d = DefaultsModeStandard
		ok = true
	case len(v) == 0:
		*d = DefaultsModeLegacy
		ok = true
	default:
		*d = DefaultsMode(v)
	}
	return ok
}
