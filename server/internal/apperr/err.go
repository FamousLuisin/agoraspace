package apperr

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