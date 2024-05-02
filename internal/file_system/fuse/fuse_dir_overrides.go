package fuse

import (
	"context"
	"fmt"
	"os"
	"path"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// Dir represents a directory.
type Dir struct {
	Path string
}

// Attr sets the attributes of the directory.
func (d *Dir) Attr(ctx context.Context, a *fuse.Attr) error {
	fi, err := os.Stat(d.Path)
	if err != nil {
		return err
	}
	a.Inode = uint64(fi.ModTime().UnixNano())
	a.Mode = fi.Mode() | os.ModeDir
	a.Size = uint64(fi.Size())
	return nil
}

// ReadDirAll lists all files in the directory.
func (d *Dir) ReadDirAll(ctx context.Context) ([]fuse.Dirent, error) {
	var dirs []fuse.Dirent
	fmt.Println("Reading directory: ")
	entries, err := os.ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		var dirType fuse.DirentType
		if entry.IsDir() {
			dirType = fuse.DT_Dir
		} else {
			dirType = fuse.DT_File
		}
		dirs = append(dirs, fuse.Dirent{
			Name: entry.Name(),
			Type: dirType,
		})
	}

	return dirs, nil
}

func (d *Dir) Lookup(ctx context.Context, name string) (fs.Node, error) {
	fmt.Println("Looking up: ", name)
	fullPath := path.Join(d.Path, name)

	// Check if this path represents a directory or a file
	fi, err := os.Stat(fullPath)
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		return &Dir{Path: fullPath}, nil
	} else {
		return &File{Path: fullPath}, nil
	}
}
