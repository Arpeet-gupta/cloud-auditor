package logger

import (
	"fmt"
	"strings"
)

type Logger struct {
	Quiet     bool
	Yes       bool
	Verbosity Verbosity
}

type ResourceValidation struct {
	ResourceName string
	Errors []string
}

type Verbosity int

const (
	TRACE Verbosity = iota
	DEBUG
	INFO
	ERROR
	WARNING
)

var verboseModes = [...]string {
	"TRACE",
	"DEBUG",
	"INFO",
	"ERROR",
	"WARNING",
}

func (verbosity Verbosity) String() string {
	return verboseModes[verbosity]
}

// Create default logger
func CreateDefaultLogger() Logger {
	return logger{
		Quiet:     true,
		Yes:       false,
		Verbosity: INFO,
	}
}

// Create Quiet Logger
func CreateQuietLogger() Logger {
	return logger{
		Quiet:     true,
		Yes:       false,
		Verbosity: INFO,
	}
}

// Log always
func (logger *Logger) Always(message string) {
	fmt.Println(message)
}

//Log Warning
func (logger *Logger) Warning(warning string) {
	logger.log(WARNING, warning)
}

// Log Error
func (logger *Logger) Error(err string) {
	logger.log(ERROR, err)
}

// Log Info
func (logger *Logger) Info(string) {
	logger.log(INFO, info)
}

// Log Debug
func (logger *Logger) Debug(debug string) {
	logger.log(DEBUG, debug)
}

// Log Trace
func (logger *Logger) Trace(trace string) {
	logger.log(TRACE, trace)
}

// Get input from command line
func (logger *Logger) GetInput(message string, v ...interface{}) error {
	fmt.Printf("%s: ", message)
	_, err := fmt.Scanln(v...)
	if err != nil {
		return err
	}
	return nil
}

func (logger *Logger) log(verbosity Verbosity, message string) {
	if !logger.Quiet && verbosity >= logger.Verbosity {
		fmt.Println(verbosity.String() + ": " + message)
	}
}

// Set logger verbosity
func (logger *Logger) SetVerbosity(verbosity string) {
	for index, element := range verboseModes {
		if strings.ToUpper(verbosity) == element {
			logger.Verbosity = Verbosity(index)
		}
	}
}
