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
