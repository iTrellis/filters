// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "crypto/sha1"

// NewSHA1 get sha1 hash repo
func NewSHA1() Repo {
	return &defaultHash{
		Hash: sha1.New(),
	}
}
