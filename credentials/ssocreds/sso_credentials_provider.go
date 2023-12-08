package ssocreds

// ProviderName is the name of the provider used to specify the source of
// credentials.
const ProviderName = "ssocreds"

type Options struct {

	// The user name that is assigned to the user.
	UserName string

	// The URL that points to the organization's ISPSS user
	// portal.
	StartURL string

	// The filepath the cached token will be retrieved from. If unset Provider will
	// use the startURL to determine the filepath at.
	//
	//    ~/.cybr/sso/cache/<sha1-hex-encoded-startURL>.json
	//
	// If custom cached token filepath is used, the Provider's startUrl
	// parameter will be ignored.
	CachedTokenFilepath string
}

// Provider is an AWS credential provider that retrieves temporary CYBR
// credentials by exchanging an SSO login token.
type Provider struct {
	options Options

	CachedTokenFilepath string
}
