package errors

const (
	ErrInternal    = "base.internal"
	ErrMsgInternal = "something went wrong"

	ErrBadRequest    = "base.bad_request"
	ErrMsgBadRequest = "arguments in request body must be properly supplied"

	ErrFoodNotFound    = "food.food_not_found"
	ErrMsgFoodNotFound = "food is not found"
)

type Err struct {
	Message string
	Code    string
}

// Error implements error interface
func (e *Err) Error() string {
	return e.Message
}

func New(message string, code string) *Err {
	return &Err{
		Message: message,
		Code:    code,
	}
}
