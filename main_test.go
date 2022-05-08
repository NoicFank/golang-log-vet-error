package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	cases := []struct{
		// variables
		name string
		data interface{}
		msgs []string
		expectNilError bool
		ErrMsg string
	} {
		{
			name: "nil data, with empty msg",
		},
		{
			name: "nil data, with one msg",
			msgs: []string{"msg1"},
			ErrMsg: "msg1",
		},
		{
			name: "nil data, with many msg",
			msgs: []string{"msg1", "msg2"},
			ErrMsg: "msg1 msg2",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := makeError(c.data, c.ErrMsg)
			if c.expectNilError {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, c.ErrMsg)
			}
		})
	}
}
