package gofortiadc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getErrorMessage(t *testing.T) {

	tests := []struct {
		input int
		want  string
	}{
		{input: -90, want: "This port is uneditable"},
		{input: -5000, want: "The traffic log size is not between 0~10 percent total hard disk size"},
		{input: 13, want: "no error message found for code 13"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, getErrorMessage(tc.input))
	}
}
