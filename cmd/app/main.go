package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adamromyn/cerebriumfuse/internal/file_system"
	cerebriumfuse "github.com/adamromyn/cerebriumfuse/internal/file_system/fuse"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// main function takes in file paths and mounts the file system
func main() {
	fmt.Println("Running")
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <nfs_dir> <cache_dir> <mount_point>", os.Args[0])
	}

	nfsDir := os.Args[1]
	cacheDir := os.Args[2]
	mountPoint := os.Args[3]

	filesystem.CacheDir = cacheDir

	c, err := fuse.Mount(
		mountPoint,
		fuse.FSName("GoFuseFS"),
		fuse.Subtype("fs"),
		fuse.AllowOther(),
	)
	if err != nil {
		log.Fatalf("Failed to mount: %v", err)
	}
	defer c.Close()

	filesys := &cerebriumfuse.RootFS{Path: nfsDir}
	if err := fs.Serve(c, filesys); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
