package kalugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		firstArg       int
		secondArg      int
		expectedErr    error
		expectedResult any
	}{
		{firstArg: 1, secondArg: 2, expectedResult: 3},
		{firstArg: 55, secondArg: 21, expectedResult: 76},
	}

	for _, tc := range testcases {
		assert := assert.New(t)

		got := Sum(tc.firstArg, tc.secondArg)

		assert.Equal(got, tc.expectedResult)

	}
}
