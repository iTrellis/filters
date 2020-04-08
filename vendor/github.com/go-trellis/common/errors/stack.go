// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package errors

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/go-trellis/common/data-structures/stack"
)

func callersDeepth(deepth int) stack.Stack {
	s := stack.New()
	for i := 0; i < deepth; i++ {
		c := caller(i + 2)
		if c.File != "" {
			s.Push(c)
		}
	}
	return s
}

func frameToString(v interface{}) string {
	if v == nil {
		return ""
	}
	fs, ok := v.([]Frame)
	if !ok {
		return ""
	}
	str := ""
	for _, f := range fs {
		str += f.ToString() + "\n"
	}
	return str
}

// Frame identifies a file, line & function name in the stack.
type Frame struct {
	File string
	Line int
	Name string
}

// ToString frame to string value
func (p *Frame) ToString() string {
	return fmt.Sprintf("%s:%d %s", p.File, p.Line, p.Name)
}

func caller(skip int) Frame {
	pc, file, line, _ := runtime.Caller(skip + 1)
	fun := runtime.FuncForPC(pc)
	return Frame{
		File: stripGOPATH(file),
		Line: line,
		Name: stripPackage(fun.Name()),
	}
}

var gopaths []string

func stripGOPATH(f string) string {
	for _, p := range gopaths {
		if strings.HasPrefix(f, p) {
			return f[len(p):]
		}
	}
	return f
}

func stripPackage(n string) string {
	slashI := strings.LastIndex(n, "/")
	if slashI == -1 {
		slashI = 0 // for built-in packages
	}
	dotI := strings.Index(n[slashI:], ".")
	if dotI == -1 {
		return n
	}
	return n[slashI+dotI+1:]
}
