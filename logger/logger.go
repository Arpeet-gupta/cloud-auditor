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
