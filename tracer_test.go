package tracer

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestTracer_New(t *testing.T) {

	tracer := New(LevelDebug)

	assert.NotNil(t, tracer)

}

func TestTracer_Level(t *testing.T) {
	tracer := New(LevelCritical)
	assert.Equal(t, LevelCritical, tracer.Level())

	tracer = New(LevelDebug)
	assert.Equal(t, LevelDebug, tracer.Level())

	tracer = New(LevelInfo)
	assert.Equal(t, LevelInfo, tracer.Level())
}

func TestTracer_TraceLevels(t *testing.T) {

	tracer := New(LevelCritical)

	tracer.Trace(LevelDebug, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelInfo, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelWarning, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelError, "%s", "test")
	assert.NotEqual(t, len(tracer.data), 1)

	tracer.Trace(LevelCritical, "%s", "test")
	assert.Equal(t, len(tracer.data), 1)

	tracer = New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")
	assert.Equal(t, len(tracer.data), 1)

	tracer.Trace(LevelInfo, "%s", "test")
	assert.Equal(t, len(tracer.data), 2)

	tracer.Trace(LevelWarning, "%s", "test")
	assert.Equal(t, len(tracer.data), 3)

	tracer.Trace(LevelError, "%s", "test")
	assert.Equal(t, len(tracer.data), 4)

	tracer.Trace(LevelCritical, "%s", "test")
	assert.Equal(t, len(tracer.data), 5)

}

func TestTracer_Trace(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")

	assert.Equal(t, tracer.data[0].data, "test")
	assert.Equal(t, tracer.data[0].level, LevelDebug)

}

func TestTracer_TraceWithInvalidLevels(t *testing.T) {

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelEverything, "%s", "test")

	}, "Trace LevelEverything")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelNothing, "%s", "test")

	}, "Trace LevelNothing")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelEverything-1, "%s", "test")

	}, "Trace too low")

	assert.Panics(t, func() {

		tracer := New(LevelDebug)

		tracer.Trace(LevelNothing+1, "%s", "test")

	}, "Trace too high")

}

func TestTracer_Copy(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "test")

	temp := tracer.Data()

	assert.Equal(t, temp[0].data, "test")

}

func TestTracer_Filter(t *testing.T) {

	tracer := New(LevelDebug)

	tracer.Trace(LevelDebug, "%s", "debug")
	tracer.Trace(LevelCritical, "%s", "critical")

	temp := tracer.Filter(LevelCritical)

	assert.Equal(t, temp[0].data, "critical")

}
