package processors

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	sdk "github.com/identityOrg/oidcsdk"
	"github.com/identityOrg/oidcsdk/impl/sdkerror"
)

type DefaultPKCEValidator struct {
	IsPKCEPlainEnabled bool
}

func NewDefaultPKCEValidator(config *sdk.Config) *DefaultPKCEValidator {
	return &DefaultPKCEValidator{IsPKCEPlainEnabled: config.PKCEPlainEnabled}
}

var b64 = base64.URLEncoding.WithPadding(base64.NoPadding)

func (d *DefaultPKCEValidator) HandleTokenEP(_ context.Context, requestContext sdk.ITokenRequestContext) sdk.IError {
	if requestContext.GetGrantType() == sdk.GrantAuthorizationCode {
		profile := requestContext.GetProfile()
		codeChallengeMethod := profile.GetCodeChallengeMethod()
		codeVerifier := requestContext.GetForm().Get("code_verifier")
		computedChallengeCode := codeVerifier
		switch codeChallengeMethod {
		case "S256":
			sum256 := sha256.Sum256([]byte(codeVerifier))
			computedChallengeCode = b64.EncodeToString(sum256[:])
			fallthrough
		case "plain":
			if len(codeVerifier) < 42 || len(codeVerifier) > 128 {
				return sdkerror.ErrInvalidGrant.WithHint("invalid entropy for code verifier")
			}
		default:
			return nil
		}
		if profile.GetCodeChallenge() != computedChallengeCode {
			return sdkerror.ErrInvalidGrant.WithHint("PKCE verification failed")
		}
	}
	return nil
}

func (d *DefaultPKCEValidator) HandleAuthEP(_ context.Context, requestContext sdk.IAuthenticationRequestContext) sdk.IError {
	form := requestContext.GetForm()
	codeChallenge := form.Get("code_challenge")
	codeChallengeMethod := form.Get("code_challenge_method")
	if codeChallenge != "" && codeChallengeMethod != "" {
		if !d.IsPKCEPlainEnabled && codeChallengeMethod == "plain" {
			return sdkerror.ErrInvalidRequest.WithHintf("un-supported code challenge method %s", codeChallengeMethod)
		}
		if codeChallengeMethod != "S256" && codeChallengeMethod != "plain" {
			return sdkerror.ErrInvalidRequest.WithHintf("un-supported code challenge method %s", codeChallengeMethod)
		}
		profile := requestContext.GetProfile()
		profile.SetCodeChallenge(codeChallenge)
		profile.SetCodeChallengeMethod(codeChallengeMethod)
	}
	return nil
}
