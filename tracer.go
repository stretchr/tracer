package tracer

import (
	"fmt"
	"github.com/stretchrcom/stew/strings"
	"time"
)

// Trace is a type that contains an actual trace
type Trace struct {
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
	data []Trace
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
	tracer.data = make([]Trace, 0, 100)

	return tracer

}

// Level gets the current level of this Tracer.
func (t *Tracer) Level() int {

	if t == nil {
		return LevelNothing
	}

	return t.level
}

// Trace saves a piece of trace data at the current time.
func (t *Tracer) Trace(level int, format string, args ...interface{}) {

	if level >= t.level && level < LevelNothing {
		trace := Trace{fmt.Sprintf(format, args...), level, time.Now()}
		t.data = append(t.data, trace)
	} else if level <= LevelEverything {
		panic("tracer: level is invalid: Cannot Trace with LevelEverything or below.")
	} else if level >= LevelNothing {
		panic("tracer: level is invalid: Cannot Trace with LevelNothing or above.")
	}

}

// Returns a copy of the trace data
func (t *Tracer) Data() []Trace {

	copiedTraces := make([]Trace, len(t.data))

	copy(copiedTraces, t.data)

	return copiedTraces

}

// Returns a copy of the trace data as an array of string
func (t *Tracer) StringData() []string {

	stringTraces := make([]string, 0, len(t.data))

	for _, trace := range t.data {
		stringTraces = append(stringTraces, fmt.Sprintf("TRACE: %s\t%s\t\t%s", trace.timestamp.String(), LevelToString(trace.level), trace.data))
	}

	return stringTraces

}

// String gets a nicely formatted string of the trace data.
func (r *Tracer) String() string {
	return strings.MergeStrings("\n", strings.JoinStrings("\n", r.StringData()...))
}

// Returns a copy of the trace data, filtered by trace level
func (t *Tracer) Filter(level int) []Trace {

	filteredTraces := make([]Trace, 0, len(t.data))

	for _, trace := range t.data {

		if trace.level == level {
			filteredTraces = append(filteredTraces, trace)
		}

	}

	return filteredTraces

}

// Returns a string representation of the level
func LevelToString(level int) string {

	switch level {
	case LevelEverything:
		return "LevelEverything"
	case LevelDebug:
		return "     LevelDebug"
	case LevelInfo:
		return "      LevelInfo"
	case LevelWarning:
		return "   LevelWarning"
	case LevelError:
		return "     LevelError"
	case LevelCritical:
		return "  LevelCritical"
	case LevelNothing:
		return "   LevelNothing"
	}
	return ""
}
