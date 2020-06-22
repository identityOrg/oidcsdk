package sdkerror

import (
	sdk "oauth2-oidc-sdk"
)

type DefaultError struct {
	Code        string
	Description string
	Status      int
	Url         string
}

func (d *DefaultError) WithDescription(desc string) sdk.IError {
	clone := DefaultError{
		Code:        d.Code,
		Description: d.Description + ": " + desc,
		Status:      d.Status,
	}
	return &clone
}

func (d *DefaultError) Error() string {
	return d.Code + ":" + d.Description
}

func (d *DefaultError) GetErrorCode() string {
	return d.Code
}

func (d *DefaultError) GetDescription() string {
	return d.Description
}

func (d *DefaultError) GetStatusCode() int {
	return d.Status
}

func (d *DefaultError) GetErrorURL() string {
	return d.Url
}

func DefaultErrorFactory(status int, code string, description string) sdk.IError {
	return &DefaultError{
		Code:        code,
		Description: description,
		Status:      status,
	}
}
