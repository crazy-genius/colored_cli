package output

import (
	"fmt"
	"io"

	"github.com/crazy-genius/colored_cli/internal/pkg/color"
)

// ColoredOutput wrapper around color
type ColoredOutput struct {
	destination io.Writer
	color       color.Color
}

// Write implement io.Writer
func (o ColoredOutput) Write(p []byte) (n int, err error) {
	if o.destination == nil {
		return 0, nil
	}

	if o.color != color.NoEntered {
		if _, e := fmt.Fprintf(o.destination, string(o.color)); e != nil {
			return 0, e
		}
	}

	return o.destination.Write(p)
}

// PaintItIn switch color output
func (o *ColoredOutput) PaintItIn(color color.Color) {
	o.color = color
}

// NewColoredOutput create new output wrapper
func NewColoredOutput(destination io.Writer) *ColoredOutput {
	return &ColoredOutput{
		destination: destination,
		color:       color.NoEntered,
	}

}
