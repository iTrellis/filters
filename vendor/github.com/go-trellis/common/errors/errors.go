// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package errors

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ErrsString get errors string
func ErrsString(errs []error) string {
	var ss []string
	for _, v := range errs {
		ss = append(ss, v.Error())
	}
	return strings.Join(ss, ";")
}

// Error error define
type Error struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace"`
	Code      uint64 `json:"code"`
	Message   string `json:"message"`
}

// ErrorCode Error functions
type ErrorCode interface {
	ID() string
	Code() uint64
	Namespace() string
	Error() string
	StackTrace() string
	Context() ErrorContext
	FullError() error
	Append(err ...interface{}) ErrorCode
	WithContext(k string, v interface{}) ErrorCode
	Marshal() ([]byte, error)
}

type errorCode struct {
	err        Error
	stackTrace string
	context    map[string]interface{}
	errors     []string
}

// NewErrorCode get a new error code
func NewErrorCode(
	errorid string, code uint64, namespace string, message string,
	stackTrace string, context map[string]interface{}) ErrorCode {
	e := &errorCode{
		err:        Error{ID: errorid, Namespace: namespace, Code: code, Message: message},
		stackTrace: stackTrace,
		context:    context,
	}

	if e.context == nil {
		e.context = make(map[string]interface{})
	}

	return e
}

func (p *errorCode) Append(err ...interface{}) ErrorCode {
	if err == nil {
		return p
	}
	for _, e := range err {
		switch ev := e.(type) {
		case ErrorCode:
			{
				p.errors = append(p.errors,
					fmt.Sprintf("(%s#%d:%s) %s", ev.Namespace(), ev.Code(), ev.ID(), ev.Error()))
			}
		case error:
			{
				p.errors = append(p.errors, ev.Error())
			}
		default:
			p.errors = append(p.errors, fmt.Sprintf("%v", e))
		}
	}
	return p
}

func (p *errorCode) Code() uint64 {
	return p.err.Code
}

func (p *errorCode) Context() ErrorContext {
	return p.context
}

func (p *errorCode) Error() string {
	msg := p.err.Message
	if len(p.errors) > 0 {
		msg = msg + "; " + strings.Join(p.errors, "; ")
	}
	return msg
}

func (p *errorCode) FullError() error {
	return fmt.Errorf(strings.Join(
		append([]string{},
			fmt.Sprintf("ID: %s#%s", genErrorCodeKey(p.Namespace(), p.Code()), p.ID()),
			"Error:", p.Error(),
			"Context:", p.Context().String(),
			"StackTrace:", p.stackTrace,
		), "\n"))
}

func (p *errorCode) ID() string {
	return p.err.ID
}

func (p *errorCode) Namespace() string {
	return p.err.Namespace
}

func (p *errorCode) Marshal() (data []byte, err error) {
	return json.Marshal(Error{
		ID:        p.err.ID,
		Code:      p.err.Code,
		Message:   p.Error(),
		Namespace: p.err.Namespace,
	})
}

func (p *errorCode) StackTrace() string {
	return p.stackTrace
}

func (p *errorCode) WithContext(key string, value interface{}) ErrorCode {
	p.context[key] = value
	return p
}

// internal functions

func genErrorCodeKey(namespace string, code uint64) string {
	return fmt.Sprintf("%s:%d", namespace, code)
}
