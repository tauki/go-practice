package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

// New method creates a new instance of the Tracer
// param w is the specified io to be used to write the output
func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}

// tracer is the Tracer that writes to an io
type tracer struct {
	out io.Writer
}

// Trace writes the passed arguments to the specified io.Writer
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}


// nilTracer is a tracer that does not interact
type nilTracer struct{}

// Trace of the nilTracer does not do anything
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
