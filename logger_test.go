package log_test

import (
	"testing"
)

import (
	"bytes"
	. "github.com/bukowa/log"
)

func TestDebug(t *testing.T) {
	// Set log level to DebugLevel
	LogLevel = DebugLevel

	// Modify flags of the Logger
	LoggerDebug.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerDebug.SetOutput(writer)

	// Test log message
	Debug("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[36mDEBUG: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestDebugf(t *testing.T) {
	// Set log level to DebugLevel
	LogLevel = DebugLevel

	// Modify flags of the Logger
	LoggerDebug.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerDebug.SetOutput(writer)

	// Test log message
	Debugf("test %s", "message")

	// Check if the log message is correct
	expectedOutput := "\x1B[36mDEBUG: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestDebugln(t *testing.T) {
	// Set log level to DebugLevel
	LogLevel = DebugLevel

	// Modify flags of the Logger
	LoggerDebug.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerDebug.SetOutput(writer)

	// Test log message
	Debugln("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[36mDEBUG: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestInfo(t *testing.T) {
	// Set log level to InfoLevel
	LogLevel = InfoLevel

	// Modify flags of the Logger
	LoggerInfo.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerInfo.SetOutput(writer)

	// Test log message
	Info("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[32mINFO: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestInfof(t *testing.T) {
	// Set log level to InfoLevel
	LogLevel = InfoLevel

	// Modify flags of the Logger
	LoggerInfo.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerInfo.SetOutput(writer)

	// Test log message
	Infof("test %s", "message")

	// Check if the log message is correct
	expectedOutput := "\x1B[32mINFO: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}
