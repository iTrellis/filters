// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

// Hash32Repo hash32 functions manager
type Hash32Repo interface {
	Sum(s string) string
	SumBytes(b []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(b []byte, times uint) string
	Sum32(b []byte) (uint32, error)
}
