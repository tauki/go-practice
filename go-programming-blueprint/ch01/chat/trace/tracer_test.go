package trace

import (
	"bytes"
	"testing"
)

// TestNew tests the trace package's Trace method
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

// TestOff checks the behavior of nilTracer
func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Something")
}
