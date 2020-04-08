// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "crypto"

// Repo hash functions manager
type Repo interface {
	Sum(s string) string
	SumBytes(b []byte) string
	SumTimes(s string, times uint) string
	SumBytesTimes(b []byte, times uint) string
}

// NewHashRepo get hash repo by crypto type
func NewHashRepo(h crypto.Hash) Repo {
	if h == crypto.MD5 {
		return NewMD5()
	} else if h == crypto.SHA1 {
		return NewSHA1()
	} else if h == crypto.SHA224 {
		return NewSHA224()
	} else if h == crypto.SHA256 {
		return NewSHA256()
	} else if h == crypto.SHA384 {
		return NewSHA384()
	} else if h == crypto.SHA512 {
		return NewSHA512()
	} else if h == crypto.SHA512_224 {
		return NewSHA512_224()
	} else if h == crypto.SHA512_256 {
		return NewSHA512_256()
	}

	return nil
}
