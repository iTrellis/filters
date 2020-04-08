// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import (
	"encoding/hex"
	"hash"
)

type defaultHash struct {
	Hash hash.Hash
}

func (p *defaultHash) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defaultHash) SumBytes(data []byte) string {
	p.Hash.Reset()
	_, err := p.Hash.Write(data)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(p.Hash.Sum(nil))
}

func (p *defaultHash) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defaultHash) SumBytesTimes(b []byte, times uint) string {
	return p.SumTimes(string(b), times)
}

type defHash32 struct {
	Hash hash.Hash32
}

func (p *defHash32) Sum(s string) string {
	return p.SumBytes([]byte(s))
}

func (p *defHash32) SumBytes(data []byte) string {
	p.Hash.Reset()
	_, err := p.Hash.Write(data)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(p.Hash.Sum(nil))
}

func (p *defHash32) SumTimes(s string, times uint) string {
	if times == 0 {
		return ""
	}

	for i := 0; i < int(times); i++ {
		s = p.Sum(s)
	}
	return s
}

func (p *defHash32) SumBytesTimes(bs []byte, times uint) string {
	return p.SumTimes(string(bs), times)
}

func (p *defHash32) Sum32(b []byte) (uint32, error) {
	_, err := p.Hash.Write(b)
	if err != nil {
		return 0, err
	}
	return p.Hash.Sum32(), nil
}
