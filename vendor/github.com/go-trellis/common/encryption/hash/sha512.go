// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "crypto/sha512"

// NewSHA384 get SHA384 hash repo
func NewSHA384() Repo {
	return &defaultHash{
		Hash: sha512.New384(),
	}
}

// NewSHA512 get SHA512 hash repo
func NewSHA512() Repo {
	return &defaultHash{
		Hash: sha512.New(),
	}
}

// NewSHA512_224 get SHA512_224 hash repo
func NewSHA512_224() Repo {
	return &defaultHash{
		Hash: sha512.New512_224(),
	}
}

// NewSHA512_256 get SHA512_256 hash repo
func NewSHA512_256() Repo {
	return &defaultHash{
		Hash: sha512.New512_256(),
	}
}
