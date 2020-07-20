package authep

import (
	"errors"
	sdk "github.com/identityOrg/oidcsdk"
	"net/http"
	"net/url"
	"strconv"
)

func DefaultAuthenticationResponseWriter(requestContext sdk.IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error {
	mode := requestContext.GetResponseMode()
	switch mode {
	case sdk.ResponseModeFragment:
		form := buildSuccessResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.Fragment = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	case sdk.ResponseModeQuery:
		form := buildSuccessResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.RawQuery = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	}
	return errors.New("invalid response mode")
}

func buildSuccessResponseForm(requestContext sdk.IAuthenticationRequestContext) url.Values {
	form := url.Values{}
	tokens := requestContext.GetIssuedTokens()
	if tokens.AccessToken != "" {
		form.Add("access_token", tokens.AccessToken)
		form.Add("expires_in", strconv.FormatInt(tokens.AccessTokenExpiry.Unix(), 10))
		form.Add("type", "bearer")
	}
	if tokens.AuthorizationCode != "" {
		form.Add(sdk.ResponseTypeCode, tokens.AuthorizationCode)
	}
	if tokens.IDToken != "" {
		form.Add("id_token", tokens.IDToken)
	}
	return form
}

func DefaultRedirectErrorWriter(requestContext sdk.IAuthenticationRequestContext, w http.ResponseWriter, r *http.Request) error {
	mode := requestContext.GetResponseMode()
	switch mode {
	case sdk.ResponseModeFragment:
		form := buildErrorResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.Fragment = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	case sdk.ResponseModeQuery:
		form := buildErrorResponseForm(requestContext)
		redirectUri, err := url.Parse(requestContext.GetRedirectURI())
		if err != nil {
			return err
		}
		redirectUri.RawQuery = form.Encode()
		http.Redirect(w, r, redirectUri.String(), http.StatusFound)
		return nil
	}
	return errors.New("invalid response mode")
}

func buildErrorResponseForm(requestContext sdk.IAuthenticationRequestContext) url.Values {
	form := url.Values{}
	err := requestContext.GetError()
	form.Add("error", err.Error())
	form.Add("error_description", err.GetDescription())
	if requestContext.GetState() != "" {
		form.Add("state", requestContext.GetState())
	}
	return form
}
