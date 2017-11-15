// +build !debug

// package inkhuffer is a package that provides debugging-by-print-statement harnesses.
//
// It's not meant to be used as a logging package.
// Ideally this package should only be imported during development and debugging phases of a program.
package inkhuffer

import "log"

const DEBUG = false

// Use tells the package to use the provided logger. NO-OP
func Use(l *log.Logger) {}

// ScopeDepth returns the level of context. NO-OP
func ScopeDepth() int { return 0 }

// EnterScope enters a context. NO-OP
func EnterScope() {}

// LeaveScope leaves a context. NO-OP
func LeaveScope() {}

// Logf writes the string and the params into the logger. NO-OP
func Logf(format string, others ...interface{}) {}

// ConfLogf is like Logf but accepts a condition. NO-OP
func CondLogf(cond func() bool, format string, others ...interface{}) {}
