package goroot

import (
	"fmt"
	"io"
	"os"
)

type Method int

const (
	Bissection Method = iota
	FPM
	NewtonRhapson
	FalsePosition
	Secant
)

type fn func(x float64) float64
type printer func(io.Writer, ...interface{})

type MethodBase struct {
	k         int
	f, g      fn
	a, b      float64
	e1, e2    float64
	x0, x1    float64
	Printable bool
	printer   printer
	w         io.Writer
}

func (mb *MethodBase) Print(args ...interface{}) {
	if mb.Printable {
		mb.printer(mb.w, args...)
	}
}

type MethodOption func(*MethodBase)

func NewMethod(opts ...MethodOption) *MethodBase {
	const (
		defaultPrintable       = false
		defaultPrec            = 0.001
		defaultIterationsLimit = 20
	)

	var (
		defaultWriter  = os.Stdout
		defaultF       = func(x float64) float64 { return 1 }
		defaultG       = func(x float64) float64 { return 1 }
		defaultPrinter = func(io.Writer, ...interface{}) { fmt.Fprintln(defaultWriter, "default printer") }
	)

	m := &MethodBase{
		f:         defaultF,
		g:         defaultG,
		e1:        defaultPrec,
		e2:        defaultPrec,
		Printable: defaultPrintable,
		printer:   defaultPrinter,
		k:         defaultIterationsLimit,
		w:         defaultWriter,
	}

	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (mb *MethodBase) GetWriter() io.Writer {
	return mb.w
}

func WithFunctions(f, h fn) MethodOption {
	return func(m *MethodBase) {
		m.f, m.g = f, h
	}
}

func WithInterval(a, b float64) MethodOption {
	return func(m *MethodBase) {
		m.a, m.b = a, b
	}
}

func WithInitialValues(x0, x1 float64) MethodOption {
	return func(m *MethodBase) {
		m.x0, m.x1 = x0, x1
	}
}

func WithPrecision(e1, e2 float64) MethodOption {
	return func(m *MethodBase) {
		m.e1, m.e2 = e1, e2
	}
}

func WithIterationsLimit(k int) MethodOption {
	return func(m *MethodBase) {
		m.k = k
	}
}

func WithPrint(wr io.Writer, p printer) MethodOption {
	return func(m *MethodBase) {

		if p != nil {
			m.printer = p
		}

		if wr != nil {
			m.w = wr
		}
		m.Printable = true
	}
}
