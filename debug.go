// +build debug

package inkhuffer

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync/atomic"
)

const DEBUG = true

var tabcount uint32
var logger = log.New(os.Stderr, "", 0)
var replacement string

// caches
var prefixes = []string{
	"",
	"\t",
	"\t\t",
	"\t\t\t",
	"\t\t\t\t",
	"\t\t\t\t\t",
}
var replacements = []string{
	"\n",
	"\n\t",
	"\n\t\t",
	"\n\t\t\t",
	"\n\t\t\t\t",
	"\n\t\t\t\t\t",
}

func Use(l *log.Logger) {
	logger = l
}

func ScopeDepth() int {
	return int(atomic.LoadUint32(&tabcount))
}

func EnterScope() {
	atomic.AddUint32(&tabcount, 1)
	tc := ScopeDepth()
	logger.SetPrefix(getPrefix(tc))
	replacement = getReplacement(tc)
}

func LeaveScope() {
	tc := ScopeDepth()
	tc--
	if tc < 0 {
		atomic.StoreUint32(&tabcount, 0)
		tc = 0
	} else {
		atomic.StoreUint32(&tabcount, uint32(tc))
	}
	logger.SetPrefix(getPrefix(tc))
	replacement = getReplacement(tc)
}

func Printf(format string, others ...interface{}) {
	s := fmt.Sprintf(format, others...)
	s = strings.Replace(s, "\n", replacement, -1)
	logger.Println(s)
}

func Logf(format string, others ...interface{}) {
	s := fmt.Sprintf(format, others...)
	s = strings.Replace(s, "\n", replacement, -1)
	logger.Println(s)
}

func CondLogf(cond func() bool, format string, others ...interface{}) {
	if cond() {
		s := fmt.Sprintf(format, others...)
		s = strings.Replace(s, "\n", replacement, -1)
		logger.Println(s)
	}
}

func getPrefix(c int) string {
	l := len(prefixes)
	if c >= l {
		diff := l - c + 1
		prefixes = append(prefixes, make([]string, diff)...)
		prefixes[c] = strings.Repeat("\t", c)
	}
	return prefixes[c]
}

func getReplacement(c int) string {
	l := len(replacements)
	if c >= l {
		diff := l - c + 1
		replacements = append(replacements, make([]string, diff)...)
		replacements[c] = "\n" + strings.Repeat("\t", c)
	}
	return replacements[c]
}
