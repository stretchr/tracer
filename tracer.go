package tracer

import (
	"fmt"
	"time"
)

// trace is a type that contains an actual trace
type trace struct {
	// data stores the trace string
	data string
	// level stores the level of this trace
	level int
	// timestamp stores the time at which this trace occurred
	timestamp time.Time
}

// tracer is a type providing trace capabilities to a go application
type Tracer struct {
	// data stores the individual string entries in the tracer
	data []trace
	// level determines what level a trace must be to actually be logged
	level int
}

const (
	// LevelEverything defines a trace level that is guarenteed to trace everything.  This trace level
	// should never be passed to the Trace command.
	LevelEverything = iota
	// LevelDebug defines a trace level that should be used for verbose tracing.
	LevelDebug
	// LevelInfo defines a trace level that should be used for normal activity tracing.
	LevelInfo
	// LevelWarning defines a trace level that should be used for tracing warnings.
	LevelWarning
	// LevelError defines a trace level that should be used for error tracing.
	LevelError
	// LevelCritical defines a trace level that should be used for critical error tracing.
	LevelCritical
	// LevelNothing defines a trace level that is guarenteed to trace nothing.  This trace level
	// should never be passed to the Trace command.
	LevelNothing
)

// New creates a new tracer
//
// The level argument is used to filter the trace to the desired level of detail.
// For example, a trace level of LevelEverything will log everything, where a trace level of LevelWarning
// will log only warnings, errors and criticals.
func New(level int) *Tracer {

	tracer := new(Tracer)

	tracer.level = level
	tracer.data = make([]trace, 0, 100)

	return tracer

}

// Trace saves a piece of trace data at the current time.
func (t *Tracer) Trace(level int, format string, args ...interface{}) {

	if level >= t.level && level < LevelNothing {
		trace := trace{fmt.Sprintf(format, args...), level, time.Now()}
		t.data = append(t.data, trace)
	} else if level <= LevelEverything {
		panic("tracer: level is invalid: Cannot Trace with LevelEverything or below.")
	} else if level >= LevelNothing {
		panic("tracer: level is invalid: Cannot Trace with LevelNothing or above.")
	}

}

// Returns a copy of the trace data
func (t *Tracer) Data() []trace {

	copiedTraces := make([]trace, len(t.data))

	copy(copiedTraces, t.data)

	return copiedTraces

}

// Returns a copy of the trace data, filtered by trace level
func (t *Tracer) Filter(level int) []trace {

	filteredTraces := make([]trace, 0, len(t.data))

	for _, trace := range t.data {

		if trace.level == level {
			filteredTraces = append(filteredTraces, trace)
		}

	}

	return filteredTraces

}
