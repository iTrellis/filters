// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package stack

import (
	"sync"
	"sync/atomic"
)

// Stack functions for manager datas in stack
type Stack interface {
	// push a data into stack
	Push(v interface{})
	// pop last data
	Pop() (interface{}, bool)
	// pop many of data
	PopMany(count int64) ([]interface{}, bool)
	// pop all data
	PopAll() ([]interface{}, bool)
	// peek last data
	Peek() (interface{}, bool)
	// get length of stack
	Length() int64
	// judge stack's lenght if 0
	IsEmpty() bool
}

type defaultStack struct {
	sync.Mutex
	length int64
	stack  []interface{}
}

// New get stack functions manager
func New() Stack {
	return &defaultStack{}
}

func (p *defaultStack) Push(v interface{}) {
	p.Lock()
	defer p.Unlock()

	prepend := make([]interface{}, 1)
	prepend[0] = v

	p.stack = append(prepend, p.stack...)
	p.length++
}

func (p *defaultStack) Pop() (v interface{}, exist bool) {
	if p.IsEmpty() {
		return
	}

	p.Lock()
	defer p.Unlock()

	v, p.stack, exist = p.stack[0], p.stack[1:], true
	p.length--

	return
}

func (p *defaultStack) PopMany(count int64) (vs []interface{}, exist bool) {

	if p.IsEmpty() {
		return
	}

	p.Lock()
	defer p.Unlock()

	if count >= p.length {
		count = p.length
	}
	p.length -= count

	vs, p.stack, exist = p.stack[:count-1], p.stack[count:], true
	return
}

func (p *defaultStack) PopAll() (all []interface{}, exist bool) {
	if p.IsEmpty() {
		return
	}
	p.Lock()
	defer p.Unlock()

	all, p.stack, exist = p.stack[:], nil, true
	p.length = 0
	return
}

func (p *defaultStack) Peek() (v interface{}, exist bool) {
	if p.IsEmpty() {
		return
	}

	p.Lock()
	defer p.Unlock()

	return p.stack[0], true
}

func (p *defaultStack) Length() int64 {
	return atomic.LoadInt64(&p.length)
}

func (p *defaultStack) IsEmpty() bool {
	return p.Length() == 0
}
