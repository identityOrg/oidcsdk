package sdkerror

var (
	InvalidRequest = DefaultErrorFactory(400, "invalid_request", "Request is invalid")
)
