package lavrushko_test

import (
	"calc/lavrushko"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	testcases := []struct {
		firstArg       int
		secondArg      int
		expectedResult int
	}{
		{firstArg: 1, secondArg: -99, expectedResult: 1},
		{firstArg: 2, secondArg: 3, expectedResult: 8},
	}

	for _, tc := range testcases {
		assert := assert.New(t)

		got := lavrushko.Pow(tc.firstArg, tc.secondArg)

		assert.Equal(got, tc.expectedResult)

	}
}
