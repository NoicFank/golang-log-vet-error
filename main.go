package main

import "fmt"

type ErrExample struct {
	data interface{}
	msg  string
}

func (e *ErrExample) Error() string {
	return e.msg
}

func makeError(data interface{}, msgAndArgs ...interface{}) error {
	msg := ""
	switch len(msgAndArgs) {
	case 0:
		// Ignore
	case 1:
		msg = fmt.Sprint(msgAndArgs[0])
	default:
		if str, ok := msgAndArgs[0].(string); ok {
			msg = fmt.Sprintf(str, msgAndArgs[1:]...)
		}
		msg = fmt.Sprint(msgAndArgs...)
	}
	return &ErrExample{data: data, msg: msg}
}

func main() {
	fmt.Println(makeError(nil, "Cannot divide %v by 0", "test"))
}
