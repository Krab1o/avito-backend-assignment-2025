package api

const (
	//Body fields
	FieldError = "errors"
	FieldToken = "token"
	//Param fields
	ParamBuying = "item"

	//400 errors
	ErrorCredsEmpty = "Credentials cannot be empty"
	ErrorBadRequest = "Bad request specified"
	ErrorNoParamSpecified = "No parameters specified"
	ErrorParametersNotAllowed = "Query parameters not allowed"
	ErrorBodyNotAllowed = "Body not allowed"
	//401 errors
	ErrorUnauthorized = "No allowance to this resource"
	ErrorInvalidAuthorizationHeader = "Invalid Authorization header format"
	ErrorToken = "Invalid or expired token"
	//500 errors
	ErrorInternalServerError = "Internal server error"
	ErrorUnknown = "Unknown error"
)
