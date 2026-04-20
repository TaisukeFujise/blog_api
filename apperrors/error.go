package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json: "-"`
}

func (myErr *MyAppError) Error() string {
	if myErr == nil {
		return "<nil>"
	}
	if myErr.Err != nil {
		return myErr.Err.Error()
	}
	return myErr.Message
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
