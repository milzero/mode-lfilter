package errors

import "fmt"

type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error: %s", err.Code, err.Message)
}

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal Error"}
)
