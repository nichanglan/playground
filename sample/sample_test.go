package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "small"},
	{50, "big"},
	{0, "zero"},
	{900, "huge"},
	{10000, "enormous"},
}

func TestSize(t *testing.T) {
	for _, test := range tests {
		size := size(test.in)
		assert.Equal(t, size, test.out)
	}
}
