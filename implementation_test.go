package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	res, err := PostfixToPrefix("1 2 + 3 4 + *")
	if assert.Nil(t, err) {
		assert.Equal(t, "*+12+3", res)
	}
}

func TestSecond(t *testing.T) {
	res, err := PostfixToPrefix("7 2 ^ 4 9 - 12 * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+^72*-4912", res)
	}
}

func TestThird(t *testing.T) {
	res, err := PostfixToPrefix("12 6 4 + 2 ^ - 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+-12^+6425", res)
	}
}

func TestFourth(t *testing.T) {
	res, err := PostfixToPrefix("x y ^ 5 s * / 10 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+/^xy*5s10", res)
	}
}

func TestFifth(t *testing.T) {
	res, err := PostfixToPrefix("x y ^ 5 s * / 10 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+/^xy*5s10", res)
	}
}

func TestSix(t *testing.T) {
	res, err := PostfixToPrefix("4 3 ^ 6 11 - 2 ^ - 1 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "+-^43^-61121", res)
	}
}

func TestErrorFirst(t *testing.T) {
	_, err := PostfixToPrefix("$ 1")
	assert.NotNil(t, err)
}

func TestErrorSecond(t *testing.T) {
	_, err := PostfixToPrefix("- - + 3 - 4")
	assert.NotNil(t, err)
}

func TestErrorThird(t *testing.T) {
	_, err := PostfixToPrefix(" ")
	assert.NotNil(t, err)
}

func TestErrorFourth(t *testing.T) {
	_, err := PostfixToPrefix("1 2 3 4 5 6 7 8")
	assert.NotNil(t, err)
}
