// GNU GPL v3 License
// Copyright (c) 2016 github.com:go-trellis

package hash

import "hash/crc32"

// NewCRC32 get crc32 hash32repo
func NewCRC32(tab *crc32.Table) Hash32Repo {
	return &defHash32{
		Hash: crc32.New(tab),
	}
}

// NewCRCIEEE get ieee hash32repo
func NewCRCIEEE() Hash32Repo {
	return &defHash32{
		Hash: crc32.NewIEEE(),
	}
}
