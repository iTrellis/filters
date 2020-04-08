// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "crypto/sha256"

// NewSHA224 get SHA224 hash repo
func NewSHA224() Repo {
	return &defaultHash{
		Hash: sha256.New224(),
	}
}

// NewSHA256 get SHA256 hash repo
func NewSHA256() Repo {
	return &defaultHash{
		Hash: sha256.New(),
	}
}
