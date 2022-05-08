package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Println(makeError(nil, "Cannot divide %v by 0", "test"))
}
