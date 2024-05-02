package fuse

import (
	"context"
	"fmt"
	"github.com/adamromyn/cerebriumfuse/internal/file_system"
	"io/ioutil"
	"os"
	"time"

	"bazil.org/fuse"
)

// File represents a file node.
type File struct {
	Path string
}

// Attr sets the attributes of a file.
func (f *File) Attr(ctx context.Context, a *fuse.Attr) error {
	fmt.Println("Calling File Attribute", f)
	fi, err := os.Stat(f.Path)
	if err != nil {
		fmt.Println("Error reading file attribute: ", err)
		return err
	}
	a.Inode = uint64(fi.ModTime().UnixNano())
	a.Mode = fi.Mode()
	a.Size = uint64(fi.Size())
	return nil
}

// ReadAll reads and returns the content of a file.
func (f *File) ReadAll(ctx context.Context) ([]byte, error) {
	fmt.Println("Reading File Content", f.Path)
	fileHash, err := filesystem.HashFileContent(f.Path)
	if err != nil {
		return nil, err
	}
	fmt.Println("File Hash: ", fileHash)

	filePath, foundFileInCache := filesystem.GetFileFromCache(fileHash)
	if !foundFileInCache {
		filePath = f.Path
		fmt.Println("File not found in cache, reading from nfs")
		time.Sleep(500 * time.Millisecond)
		go filesystem.AddFileToCache(filePath, fileHash)
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil

}
