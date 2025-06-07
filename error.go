package kiwivm

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("error: [%d], message: [%s]", e.Code, e.Message)
}

var UnknownError = &Error{
	Code:    -1,
	Message: "unknown error",
}
