package api

const (
	//Body fields
	FieldError = "errors"
	FieldToken = "token"
	//Param fields
	ParamBuying = "item"

	//400 errors
	ErrorBadRequest = "Bad request specified"
	ErrorNoParamSpecified = "No parameters specified"
	ErrorParametersNotAllowed = "Query parameters are not allowed"
	//401 errors
	ErrorUnauthorized = "No allowance to this resource"
	ErrorInvalidAuthorizationHeader = "Invalid Authorization header format"
	ErrorToken = "Invalid or expired token"
	//500 errors
	ErrorInternalServerError = "Internal server error"
)
