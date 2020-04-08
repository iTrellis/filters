// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "crypto/md5"

// NewMD5 get md5 hash repo
func NewMD5() Repo {
	return &defaultHash{
		Hash: md5.New(),
	}
}
