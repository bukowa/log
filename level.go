package log

import (
	"strconv"
	"strings"
)

var (
	// LogLevel is the global log level that controls which log messages are printed.
	LogLevel = InfoLevel
)

const (
	// DebugLevel is the lowest severity level.
	DebugLevel Level = iota
	// InfoLevel is the second-lowest severity level.
	InfoLevel
	// WarningLevel is the second-highest severity level.
	WarningLevel
	// ErrorLevel is the third-highest severity level.
	ErrorLevel
	// FatalLevel is the highest severity level.
	FatalLevel
)

// LevelNames is a mapping of Level values to their corresponding string names.
var LevelNames = map[Level]string{
	DebugLevel:   "DEBUG",
	InfoLevel:    "INFO",
	WarningLevel: "WARNING",
	ErrorLevel:   "ERROR",
	FatalLevel:   "CRITICAL",
}

// Level is an integer type representing the severity level of a log message.
type Level int32

// String returns a string representation of the Level value.
func (l *Level) String() string {
	return LevelNames[*l]
}

// Set sets the value of the Level to the integer representation of the input string.
// If the input string matches one of the strings in the LevelNames map, the corresponding Level value is set.
// Otherwise, the Level value is set to the integer representation of the input string.
// If the input string cannot be converted to an integer, an error is returned.
func (l *Level) Set(s string) error {
	for k, v := range LevelNames {
		if strings.ToLower(s) == strings.ToLower(v) {
			*l = k
		}
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*l = Level(n)
	return err
}

// IsLogged returns true if the desired log level is greater than or equal to the current log level.
// This is used to determine if a log message should be printed or not.
func IsLogged(desired, level Level) bool {
	if level >= desired {
		return true
	}
	return false
}
