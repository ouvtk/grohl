package grohl

import (
	"bytes"
	"testing"
	"time"
)

func TestLogsCounter(t *testing.T) {
	s, buf := statterWithBuffer()
	s.Counter(1.0, "a", 1, 2)
	if line := logged(buf); line != "statsd=a count=1\nstatsd=a count=2" {
		t.Errorf("Bad log output: %q", line)
	}
}

func TestLogsTiming(t *testing.T) {
	s, buf := statterWithBuffer()
	dur1, _ := time.ParseDuration("15ms")
	dur2, _ := time.ParseDuration("3s")

	s.Timing(1.0, "a", dur1, dur2)
	if line := logged(buf); line != "statsd=a timing=15\nstatsd=a timing=3000" {
		t.Errorf("Bad log output: %q", line)
	}
}

func TestLogsGauge(t *testing.T) {
	s, buf := statterWithBuffer()
	s.Gauge(1.0, "a", "1", "2")
	if line := logged(buf); line != "statsd=a gauge=1\nstatsd=a gauge=2" {
		t.Errorf("Bad log output: %q", line)
	}
}

func TestDefaultLogger(t *testing.T) {
	logger, buf := loggerWithBuffer()
	SetLogger(logger)
	s := NewStatter(nil)
	if s.Logger != logger {
		t.Errorf("Wrong logger: ", s.Logger)
	}

	s.Counter(1.0, "a", 1)

	if len(buf.String()) == 0 {
		t.Errorf("nothing written to the buffer")
	}
}

func statterWithBuffer() (*Statter, *bytes.Buffer) {
	logger, buf := loggerWithBuffer()
	return NewStatter(logger), buf
}