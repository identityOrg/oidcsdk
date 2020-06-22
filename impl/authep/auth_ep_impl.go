package authep

import (
	"net/url"
	"oauth2-oidc-sdk"
)

type (
	DefaultAuthenticationRequest struct {
		RequestID         string
		State             string
		ClientID          string
		RedirectURI       string
		ResponseType      oauth2_oidc_sdk.Arguments
		Nonce             string
		RequestedScopes   oauth2_oidc_sdk.Arguments
		RequestedAudience oauth2_oidc_sdk.Arguments
		Form              *url.Values
	}
	DefaultAuthenticationResponse struct {
		State        string
		RedirectURI  string
		ResponseMode string
	}
)
