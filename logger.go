package log

import (
	"log"
	"os"
)

var (
	// LoggerDebug is the logger for debug level log messages.
	LoggerDebug = log.New(os.Stdout, "\x1B[36mDEBUG: \x1B[0m", log.Ldate|log.Ltime|log.Lmsgprefix)
	// LoggerInfo is the logger for info level log messages.
	LoggerInfo = log.New(os.Stdout, "\x1B[32mINFO: \x1B[0m", log.Ldate|log.Ltime|log.Lmsgprefix)
	// LoggerWarning is the logger for warning level log messages.
	LoggerWarning = log.New(os.Stdout, "\x1B[35mWARN: \x1B[0m", log.Ldate|log.Ltime|log.Lmsgprefix)
	// LoggerError is the logger for error level log messages.
	LoggerError = log.New(os.Stdout, "\x1B[91mERROR: \x1B[0m", log.Ldate|log.Ltime|log.Lmsgprefix)
	// LoggerFatal is the logger for fatal level log messages.
	LoggerFatal = log.New(os.Stderr, "\x1B[31mFATAL: \x1B[0m", log.Ldate|log.Ltime|log.Lmsgprefix)
)

// Debug logs a message with the DebugLevel severity level.
// If the desired log level (LogLevel) is equal to or greater than the DebugLevel, the message will be logged.
func Debug(a ...interface{}) {
	if IsLogged(LogLevel, DebugLevel) {
		LoggerDebug.Print(a...)
	}
}

// Debugf logs a message with the DebugLevel severity level using the provided format string.
// If the desired log level (LogLevel) is equal to or greater than the DebugLevel, the message will be logged.
func Debugf(format string, a ...interface{}) {
	if IsLogged(LogLevel, DebugLevel) {
		LoggerDebug.Printf(format, a...)
	}
}

// Debugln logs a message with the DebugLevel severity level and appends a newline.
// If the desired log level (LogLevel) is equal to or greater than the DebugLevel, the message will be logged.
func Debugln(a ...interface{}) {
	if IsLogged(LogLevel, DebugLevel) {
		LoggerDebug.Println(a...)
	}
}

// Info logs a message at the info severity level.
// If the desired log level (LogLevel) is equal to or greater than the InfoLevel, the message will be logged.
func Info(a ...interface{}) {
	if IsLogged(LogLevel, InfoLevel) {
		LoggerInfo.Print(a...)
	}
}

// Infof logs a formatted message at the info severity level.
// If the desired log level (LogLevel) is equal to or greater than the InfoLevel, the message will be logged.
func Infof(format string, a ...interface{}) {
	if IsLogged(LogLevel, InfoLevel) {
		LoggerInfo.Printf(format, a...)
	}
}

// Infoln logs a message at the info severity level, followed by a newline.
// If the desired log level (LogLevel) is equal to or greater than the InfoLevel, the message will be logged.
func Infoln(a ...interface{}) {
	if IsLogged(LogLevel, InfoLevel) {
		LoggerInfo.Println(a...)
	}
}

// Warning logs a message with severity WarningLevel.
// If the desired log level (LogLevel) is equal to or greater than the WarningLevel, the message will be logged.
func Warning(a ...interface{}) {
	if IsLogged(LogLevel, WarningLevel) {
		LoggerWarning.Print(a...)
	}
}

// Warningf logs a message with severity WarningLevel.
// If the desired log level (LogLevel) is equal to or greater than the WarningLevel, the message will be logged.
func Warningf(format string, a ...interface{}) {
	if IsLogged(LogLevel, WarningLevel) {
		LoggerWarning.Printf(format, a...)
	}
}

// Warningln logs a message with severity WarningLevel.
// If the desired log level (LogLevel) is equal to or greater than the WarningLevel, the message will be logged.
func Warningln(a ...interface{}) {
	if IsLogged(LogLevel, WarningLevel) {
		LoggerWarning.Println(a...)
	}
}

// Error logs a message with the ErrorLevel severity level.
// If the desired log level (LogLevel) is equal to or greater than the ErrorLevel, the message will be logged.
func Error(a ...interface{}) {
	if IsLogged(LogLevel, ErrorLevel) {
		LoggerError.Print(a...)
	}
}

// Errorf logs a message with the ErrorLevel severity level using the provided format string.
// If the desired log level (LogLevel) is equal to or greater than the ErrorLevel, the message will be logged.
func Errorf(format string, a ...interface{}) {
	if IsLogged(LogLevel, ErrorLevel) {
		LoggerError.Printf(format, a...)
	}
}

// Errorln logs a message with the ErrorLevel severity level and appends a newline.
// If the desired log level (LogLevel) is equal to or greater than the ErrorLevel, the message will be logged.
func Errorln(a ...interface{}) {
	if IsLogged(LogLevel, ErrorLevel) {
		LoggerError.Println(a...)
	}
}

// Fatal logs a message at the fatal severity level.
// If the desired log level (LogLevel) is equal to or greater than the FatalLevel, the message will be logged
// and the program will terminate.
func Fatal(a ...interface{}) {
	if IsLogged(LogLevel, FatalLevel) {
		LoggerFatal.Fatal(a...)
	}
}

// Fatalf logs a formatted message at the fatal severity level.
// If the desired log level (LogLevel) is equal to or greater than the FatalLevel, the message will be logged
// and the program will terminate.
func Fatalf(format string, a ...interface{}) {
	if IsLogged(LogLevel, FatalLevel) {
		LoggerFatal.Fatalf(format, a...)
	}
}

// Fatalln logs a message at the fatal severity level, followed by a newline.
// If the desired log level (LogLevel) is equal to or greater than the FatalLevel, the message will be logged
// and the program will terminate.
func Fatalln(a ...interface{}) {
	if IsLogged(LogLevel, FatalLevel) {
		LoggerFatal.Fatalln(a...)
	}
}
