package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	res, err := PostfixToPrefix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 2 - 3 * 5 +", res)
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
