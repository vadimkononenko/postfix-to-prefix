package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	input := strings.NewReader("")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "", output.String())
	}
}

func TestSimple(t *testing.T) {
	input := strings.NewReader("1 2 + 3 4 + *")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "*+12+34", output.String())
	}
}

func TestError(t *testing.T) {
	input := strings.NewReader("+ 2")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}
