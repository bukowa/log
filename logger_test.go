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

func TestInfoln(t *testing.T) {
	// Set log level to InfoLevel
	LogLevel = InfoLevel

	// Modify flags of the Logger
	LoggerInfo.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerInfo.SetOutput(writer)

	// Test log message
	Infoln("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[32mINFO: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestWarning(t *testing.T) {
	// Set log level to WarningLevel
	LogLevel = WarningLevel

	// Modify flags of the Logger
	LoggerWarning.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerWarning.SetOutput(writer)

	// Test log message
	Warning("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[35mWARN: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestWarningf(t *testing.T) {
	// Set log level to WarningLevel
	LogLevel = WarningLevel

	// Modify flags of the Logger
	LoggerWarning.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerWarning.SetOutput(writer)

	// Test log message
	Warningf("test %s", "message")

	// Check if the log message is correct
	expectedOutput := "\x1B[35mWARN: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestWarningln(t *testing.T) {
	// Set log level to WarningLevel
	LogLevel = WarningLevel

	// Modify flags of the Logger
	LoggerWarning.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerWarning.SetOutput(writer)

	// Test log message
	Warningln("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[35mWARN: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestError(t *testing.T) {
	// Set log level to ErrorLevel
	LogLevel = ErrorLevel

	// Modify flags of the Logger
	LoggerError.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerError.SetOutput(writer)

	// Test log message
	Error("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[91mERROR: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestErrorf(t *testing.T) {
	// Set log level to ErrorLevel
	LogLevel = ErrorLevel

	// Modify flags of the Logger
	LoggerError.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerError.SetOutput(writer)

	// Test log message
	Errorf("test message %d", 1)

	// Check if the log message is correct
	expectedOutput := "\x1B[91mERROR: \x1B[0mtest message 1\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestErrorln(t *testing.T) {
	// Set log level to ErrorLevel
	LogLevel = ErrorLevel

	// Modify flags of the Logger
	LoggerError.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerError.SetOutput(writer)

	// Test log message
	Errorln("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[91mERROR: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}

func TestFatal(t *testing.T) {
	// Set log level to FatalLevel
	LogLevel = FatalLevel

	// Modify flags of the Logger
	LoggerFatal.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerFatal.SetOutput(writer)

	// Set FatalFunc to a custom function that records if it was called
	called := false
	FatalFunc = func() { called = true }

	// Test log message
	Fatal("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[31mFATAL: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}

	// Check if FatalFunc was called
	if !called {
		t.Errorf("FatalFunc was not called")
	}
}

func TestFatalf(t *testing.T) {
	// Set log level to FatalLevel
	LogLevel = FatalLevel

	// Modify flags of the Logger
	LoggerFatal.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerFatal.SetOutput(writer)

	// Set FatalFunc to a custom function that records if it was called
	called := false
	FatalFunc = func() { called = true }

	// Test log message
	Fatalf("test %s", "message")

	// Check if the log message is correct
	expectedOutput := "\x1B[31mFATAL: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}

	// Check if FatalFunc was called
	if !called {
		t.Errorf("FatalFunc was not called")
	}
}

func TestFatalln(t *testing.T) {
	// Set log level to FatalLevel
	LogLevel = FatalLevel

	// Modify flags of the Logger
	LoggerFatal.SetFlags(0)

	// Create a custom writer to capture the log output
	writer := &bytes.Buffer{}
	LoggerFatal.SetOutput(writer)

	// Set a custom FatalFunc
	FatalFunc = func() {
		// Capture the panic message
		defer func() {
			if r := recover(); r != nil {
				t.Log(r)
			}
		}()
		// Call panic
		panic("test message")
	}

	// Test log message
	Fatalln("test message")

	// Check if the log message is correct
	expectedOutput := "\x1B[31mFATAL: \x1B[0mtest message\n"
	if writer.String() != expectedOutput {
		t.Errorf("Unexpected log output.\nExpected: %s\nActual: %s", expectedOutput, writer.String())
	}
}
