package main

import "fmt"

type ErrExample struct {
	data interface{}
	msg  string
}

func (e *ErrExample) Error() string {
	return e.msg
}

// using slice will be more comfortable
func makeError(data interface{}, msgAndArgs ...interface{}) error {
	msg := fmt.Sprintf("%v", msgAndArgs)
	return &ErrExample{data: data, msg: msg}
}

