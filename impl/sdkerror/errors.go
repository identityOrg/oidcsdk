package sdkerror

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

var (
	// ErrInvalidatedAuthorizeCode is an error indicating that an authorization code has been
	// used previously.
	ErrInvalidatedAuthorizeCode = errors.New("Authorization code has ben invalidated")
	// ErrSerializationFailure is an error indicating that the transactional capable storage could not guarantee
	// consistency of Update & Delete operations on the same rows between multiple sessions.
	ErrSerializationFailure = errors.New("The request could not be completed due to concurrent access")
	ErrUnknownRequest       = &SDKError{
		Name:        errUnknownErrorName,
		Description: "The handler is not responsible for this request",
		Code:        http.StatusBadRequest,
	}
	ErrRequestForbidden = &SDKError{
		Name:        errRequestForbidden,
		Description: "The request is not allowed",
		Hint:        "You are not allowed to perform this action.",
		Code:        http.StatusForbidden,
	}
	ErrInvalidRequest = &SDKError{
		Name:        errInvalidRequestName,
		Description: "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed",
		Hint:        "Make sure that the various parameters are correct, be aware of case sensitivity and trim your parameters. Make sure that the client you are using has exactly whitelisted the redirect_uri you specified.",
		Code:        http.StatusBadRequest,
	}
	ErrUnauthorizedClient = &SDKError{
		Name:        errUnauthorizedClientName,
		Description: "The client is not authorized to request a token using this method",
		Hint:        "Make sure that client id and secret are correctly specified and that the client exists.",
		Code:        http.StatusBadRequest,
	}
	ErrAccessDenied = &SDKError{
		Name:        errAccessDeniedName,
		Description: "The resource owner or authorization server denied the request",
		Hint:        "Make sure that the request you are making is valid. Maybe the credential or request parameters you are using are limited in scope or otherwise restricted.",
		Code:        http.StatusForbidden,
	}
	ErrUnsupportedResponseType = &SDKError{
		Name:        errUnsupportedResponseTypeName,
		Description: "The authorization server does not support obtaining a token using this method",
		Code:        http.StatusBadRequest,
	}
	ErrInvalidScope = &SDKError{
		Name:        errInvalidScopeName,
		Description: "The requested scope is invalid, unknown, or malformed",
		Code:        http.StatusBadRequest,
	}
	ErrServerError = &SDKError{
		Name:        errServerErrorName,
		Description: "The authorization server encountered an unexpected condition that prevented it from fulfilling the request",
		Code:        http.StatusInternalServerError,
	}
	ErrTemporarilyUnavailable = &SDKError{
		Name:        errTemporarilyUnavailableName,
		Description: "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server",
		Code:        http.StatusServiceUnavailable,
	}
	ErrUnsupportedGrantType = &SDKError{
		Name:        errUnsupportedGrantTypeName,
		Description: "The authorization grant type is not supported by the authorization server",
		Code:        http.StatusBadRequest,
	}
	ErrInvalidGrant = &SDKError{
		Name:        errInvalidGrantName,
		Description: "The provided authorization grant (e.g., authorization code, resource owner credentials) or refresh token is invalid, expired, revoked, does not match the redirection URI used in the authorization request, or was issued to another client",
		Code:        http.StatusBadRequest,
	}
	ErrInvalidClient = &SDKError{
		Name:        errInvalidClientName,
		Description: "Client authentication failed (e.g., unknown client, no client authentication included, or unsupported authentication method)",
		Code:        http.StatusUnauthorized,
	}
	ErrInvalidState = &SDKError{
		Name:        errInvalidStateName,
		Description: fmt.Sprintf("The state is missing or has less than required characters and is therefore considered too weak"),
		Code:        http.StatusBadRequest,
	}
	ErrMisconfiguration = &SDKError{
		Name:        errMisconfigurationName,
		Description: "The request failed because of an internal error that is probably caused by misconfiguration",
		Code:        http.StatusInternalServerError,
	}
	ErrInsufficientEntropy = &SDKError{
		Name:        errInsufficientEntropyName,
		Description: fmt.Sprintf("The request used a security parameter (e.g., anti-replay, anti-csrf) with insufficient entropy"),
		Code:        http.StatusBadRequest,
	}
	ErrNotFound = &SDKError{
		Name:        errNotFoundName,
		Description: "Could not find the requested resource(s)",
		Code:        http.StatusNotFound,
	}
	ErrRequestUnauthorized = &SDKError{
		Name:        errRequestUnauthorizedName,
		Description: "The request could not be authorized",
		Hint:        "Check that you provided valid credentials in the right format.",
		Code:        http.StatusUnauthorized,
	}
	ErrTokenSignatureMismatch = &SDKError{
		Name:        errTokenSignatureMismatchName,
		Description: "Token signature mismatch",
		Hint:        "Check that you provided  a valid token in the right format.",
		Code:        http.StatusBadRequest,
	}
	ErrInvalidTokenFormat = &SDKError{
		Name:        errInvalidTokenFormatName,
		Description: "Invalid token format",
		Hint:        "Check that you provided a valid token in the right format.",
		Code:        http.StatusBadRequest,
	}
	ErrTokenExpired = &SDKError{
		Name:        errTokenExpiredName,
		Description: "Token expired",
		Hint:        "The token expired.",
		Code:        http.StatusUnauthorized,
	}
	ErrScopeNotGranted = &SDKError{
		Name:        errScopeNotGrantedName,
		Description: "The token was not granted the requested scope",
		Hint:        "The resource owner did not grant the requested scope.",
		Code:        http.StatusForbidden,
	}
	ErrTokenClaim = &SDKError{
		Name:        errTokenClaimName,
		Description: "The token failed validation due to a claim mismatch",
		Hint:        "One or more token claims failed validation.",
		Code:        http.StatusUnauthorized,
	}
	ErrInactiveToken = &SDKError{
		Name:        errTokenInactiveName,
		Description: "Token is inactive because it is malformed, expired or otherwise invalid",
		Hint:        "Token validation failed.",
		Code:        http.StatusUnauthorized,
	}
	ErrRevokationClientMismatch = &SDKError{
		Name:        errRevokationClientMismatchName,
		Description: "Token was not issued to the client making the revokation request",
		Code:        http.StatusBadRequest,
	}
	ErrLoginRequired = &SDKError{
		Name:        errLoginRequired,
		Description: "The Authorization Server requires End-User authentication",
		Code:        http.StatusBadRequest,
	}
	ErrInteractionRequired = &SDKError{
		Description: "The Authorization Server requires End-User interaction of some form to proceed",
		Name:        errInteractionRequired,
		Code:        http.StatusBadRequest,
	}
	ErrConsentRequired = &SDKError{
		Description: "The Authorization Server requires End-User consent",
		Name:        errConsentRequired,
		Code:        http.StatusBadRequest,
	}
	ErrRequestNotSupported = &SDKError{
		Description: "The OP does not support use of the request parameter",
		Name:        errRequestNotSupportedName,
		Code:        http.StatusBadRequest,
	}
	ErrRequestURINotSupported = &SDKError{
		Description: "The OP does not support use of the request_uri parameter",
		Name:        errRequestURINotSupportedName,
		Code:        http.StatusBadRequest,
	}
	ErrRegistrationNotSupported = &SDKError{
		Description: "The OP does not support use of the registration parameter",
		Name:        errRegistrationNotSupportedName,
		Code:        http.StatusBadRequest,
	}
	ErrInvalidRequestURI = &SDKError{
		Description: "The request_uri in the Authorization Request returns an error or contains invalid data. ",
		Name:        errInvalidRequestURI,
		Code:        http.StatusBadRequest,
	}
	ErrInvalidRequestObject = &SDKError{
		Description: "The request parameter contains an invalid Request Object. ",
		Name:        errInvalidRequestObject,
		Code:        http.StatusBadRequest,
	}
	ErrJTIKnown = &SDKError{
		Description: "The jti was already used.",
		Name:        errJTIKnownName,
		Code:        http.StatusBadRequest,
	}
	ErrUnsupportedTokenType = &SDKError{
		Name:        errUnsupportedTokenTypeName,
		Description: "The authorization server does not support the revocation of the presented token type.  That is, the client tried to revoke an access token on a server not supporting this feature.",
		Code:        http.StatusBadRequest,
	}
)

const (
	errInvalidRequestURI           = "invalid_request_uri"
	errInvalidRequestObject        = "invalid_request_object"
	errConsentRequired             = "consent_required"
	errInteractionRequired         = "interaction_required"
	errLoginRequired               = "login_required"
	errRequestUnauthorizedName     = "request_unauthorized"
	errRequestForbidden            = "request_forbidden"
	errInvalidRequestName          = "invalid_request"
	errUnauthorizedClientName      = "unauthorized_client"
	errAccessDeniedName            = "access_denied"
	errUnsupportedResponseTypeName = "unsupported_response_type"
	errInvalidScopeName            = "invalid_scope"
	errServerErrorName             = "server_error"
	errTemporarilyUnavailableName  = "temporarily_unavailable"
	errUnsupportedGrantTypeName    = "unsupported_grant_type"
	errInvalidGrantName            = "invalid_grant"
	errInvalidClientName           = "invalid_client"
	errNotFoundName                = "not_found"
	errInvalidStateName            = "invalid_state"
	errMisconfigurationName        = "misconfiguration"
	errInsufficientEntropyName     = "insufficient_entropy"
	errInvalidTokenFormatName      = "invalid_token"
	errTokenSignatureMismatchName  = "token_signature_mismatch"
	errTokenExpiredName            = "token_expired"
	errScopeNotGrantedName         = "scope_not_granted"
	errTokenClaimName              = "token_claim"
	errTokenInactiveName           = "token_inactive"
	// errAuthorizaionCodeInactiveName = "authorization_code_inactive"
	errUnknownErrorName             = "error"
	errRevokationClientMismatchName = "revokation_client_mismatch"
	errRequestNotSupportedName      = "request_not_supported"
	errRequestURINotSupportedName   = "request_uri_not_supported"
	errRegistrationNotSupportedName = "registration_not_supported"
	errJTIKnownName                 = "jti_known"
	errUnsupportedTokenTypeName     = "unsupported_token_type"
)

func ErrorToSDKError(err error) *SDKError {
	if e, ok := err.(*SDKError); ok {
		return e
	} else if e, ok := errors.Cause(err).(*SDKError); ok {
		return e
	}
	return &SDKError{
		Name:        errUnknownErrorName,
		Description: "The error is unrecognizable.",
		Debug:       err.Error(),
		Code:        http.StatusInternalServerError,
	}
}

type SDKError struct {
	Name        string `json:"error"`
	Description string `json:"error_description"`
	Hint        string `json:"error_hint,omitempty"`
	Code        int    `json:"status_code,omitempty"`
	Debug       string `json:"error_debug,omitempty"`
}

func (e *SDKError) GetReason() string {
	return e.Hint
}

func (e *SDKError) GetDescription() string {
	return e.Description
}

func (e *SDKError) GetDebugInfo() string {
	return e.Debug
}

func (e *SDKError) GetStatus() string {
	return http.StatusText(e.Code)
}

func (e *SDKError) Error() string {
	return e.Name
}

func (e *SDKError) GetStatusCode() int {
	return e.Code
}

func (e *SDKError) WithHintf(hint string, args ...interface{}) *SDKError {
	return e.WithHint(fmt.Sprintf(hint, args...))
}

func (e *SDKError) WithHint(hint string) *SDKError {
	err := *e
	err.Hint = hint
	return &err
}

func (e *SDKError) WithDebug(debug string) *SDKError {
	err := *e
	err.Debug = debug
	return &err
}

func (e *SDKError) WithDebugf(debug string, args ...interface{}) *SDKError {
	return e.WithDebug(fmt.Sprintf(debug, args...))
}

func (e *SDKError) WithDescription(description string) *SDKError {
	err := *e
	err.Description = description
	return &err
}
