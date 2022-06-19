package lab2

import (
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, ch.Input)
	if err != nil {
		return err
	}

	res, err := PostfixToPrefix(buf.String())
	if err != nil {
		return err
	}

	ch.Output.Write([]byte(res))
	return nil
}
