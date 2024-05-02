package fuse

import (
	"bazil.org/fuse/fs"
)

// RootFS represents the root filesystem.
type RootFS struct {
	Path string
}

// Root returns the root directory of the filesystem.
func (fs *RootFS) Root() (fs.Node, error) {
	return &Dir{Path: fs.Path}, nil
}
