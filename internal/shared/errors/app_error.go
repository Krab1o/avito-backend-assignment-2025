package errs

type AppError struct {
	ErrType ErrorType
	Message string
	Err  	error
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func newAppError(errType ErrorType, msg string, err error) *AppError {
	return &AppError{
		ErrType: errType,
		Message: msg,
		Err:	 err,
	}
}

func NewDBError(msg string, err error) *AppError {
	return newAppError(DBError, msg, err)
}

func NewSemanticError(msg string, err error) *AppError {
	return newAppError(SemanticError, msg, err)
}

func NewServiceError(msg string, err error) *AppError {
	return newAppError(ServiceError, msg, err)
}

func NewNotFoundError(msg string, err error) *AppError {
	return newAppError(NotFound, msg, err)
}

func NewBadRequestError(msg string, err error) *AppError {
	return newAppError(BadRequest, msg, err)
}

func NewUnauthorizedError(msg string, err error) *AppError {
	return newAppError(Unauthorized, msg, err)
}
