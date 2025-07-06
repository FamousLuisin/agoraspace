package apperr

const (
	ErrBadRequest          = "BAD_REQUEST"
	ErrUnauthorized        = "UNAUTHORIZED"
	ErrForbidden           = "FORBIDDEN"
	ErrNotFound            = "NOT_FOUND"
	ErrMethodNotAllowed    = "METHOD_NOT_ALLOWED"
	ErrConflict            = "CONFLICT"
	ErrUnprocessableEntity = "UNPROCESSABLE_ENTITY"
	ErrInternalServer      = "INTERNAL_SERVER_ERROR"
	ErrServiceUnavailable  = "SERVICE_UNAVAILABLE"
)

type AppErr struct {
	Message string
	Err     string
	Code    int
}

func (e *AppErr) Error() string {
	return e.Message
}

func NewAppError(message, err string, code int) *AppErr {
	return &AppErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}