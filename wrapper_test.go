package restyzerolog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog"
	"resty.dev/v3"
)

type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func TestNew(t *testing.T) {
	var buf bytes.Buffer

	var wrappedLogger any

	logger := zerolog.New(&buf)
	wrappedLogger = New(logger)

	if _, ok := wrappedLogger.(resty.Logger); !ok {
		t.Errorf("Expected wrapped logger to implement resty.Logger interface")
	}
}

func TestLogger_Errorf(t *testing.T) {
	var buf bytes.Buffer

	msg := &message{}
	logger := zerolog.New(&buf)
	wrappedLogger := New(logger)

	wrappedLogger.Errorf("error message: %s", "test")

	if err := json.Unmarshal(buf.Bytes(), msg); err != nil {
		t.Fatalf("Expected log output to be a valid JSON, got: %s", buf.String())
	}

	if msg.Level != "error" {
		t.Fatalf("Expected log level to be 'error', got: %s", msg.Level)
	}

	if msg.Message != "error message: test" {
		t.Fatalf("Expected log message to be 'error message: test', got: %s", msg.Message)
	}
}

func TestLogger_Warnf(t *testing.T) {
	var buf bytes.Buffer

	msg := &message{}
	logger := zerolog.New(&buf)
	wrappedLogger := New(logger)

	wrappedLogger.Warnf("warn message: %s", "test")

	if err := json.Unmarshal(buf.Bytes(), msg); err != nil {
		t.Fatalf("Expected log output to be a valid JSON, got: %s", buf.String())
	}

	if msg.Level != "warn" {
		t.Fatalf("Expected log level to be 'warn', got: %s", msg.Level)
	}

	if msg.Message != "warn message: test" {
		t.Fatalf("Expected log message to be 'warn message: test', got: %s", msg.Message)
	}
}

func TestLogger_Debugf(t *testing.T) {
	var buf bytes.Buffer

	msg := &message{}
	logger := zerolog.New(&buf)
	wrappedLogger := New(logger)

	wrappedLogger.Debugf("debug message: %s", "test")

	if err := json.Unmarshal(buf.Bytes(), msg); err != nil {
		t.Fatalf("Expected log output to be a valid JSON, got: %s", buf.String())
	}

	if msg.Level != "debug" {
		t.Fatalf("Expected log level to be 'debug', got: %s", msg.Level)
	}

	if msg.Message != "debug message: test" {
		t.Fatalf("Expected log message to be 'debug message: test', got: %s", msg.Message)
	}
}
