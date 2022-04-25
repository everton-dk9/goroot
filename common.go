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
	printable bool
	printer   printer
	w         io.Writer
}

func (mb *MethodBase) IsPrintable() bool {
	return mb.printable
}

func (mb *MethodBase) Print(args ...interface{}) {
	if mb.IsPrintable() {
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
		printable: defaultPrintable,
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

func WithFunctions(f, g fn) MethodOption {
	return func(m *MethodBase) {
		if f != nil {
			m.f = f
		}

		if g != nil {
			m.g = g
		}
	}
}

func WithInterval(a, b float64) MethodOption {
	return func(m *MethodBase) {
		// if a >= 0 {
		// }

		// if b >= 0 {
		// 	m.b = b
		// }
		m.a, m.b = a, b
	}
}

func WithInitialValues(x0, x1 float64) MethodOption {
	return func(m *MethodBase) {
		if x0 >= 0 {
			m.x0 = x0
		}

		if x1 >= 0 {
			m.x1 = x1
		}
	}
}

func WithPrecision(e1, e2 float64) MethodOption {
	return func(m *MethodBase) {
		if e1 >= 0 {
			m.e1 = e1
		}

		if e2 >= 0 {
			m.e2 = e2
		}
	}
}

func WithIterationsLimit(k int) MethodOption {
	return func(m *MethodBase) {
		if k >= 0 {
			m.k = k
		}
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
		m.printable = true
	}
}
