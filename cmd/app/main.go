package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/adamromyn/cerebriumfuse/internal/file_system"
	cerebriumfuse "github.com/adamromyn/cerebriumfuse/internal/file_system/fuse"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
)

// main function takes in file paths and mounts the file system
func main() {
	addConsoleVibesOnStart()
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <nfs_dir> <cache_dir> <mount_point>", os.Args[0])
	}

	nfsDir := os.Args[1]
	cacheDir := os.Args[2]
	mountPoint := os.Args[3]

	filesystem.CacheDir = cacheDir
	clearCacheDirectory(cacheDir)
	unmountFS(mountPoint)

	fmt.Println("Mounting filesystem...")
	c, err := fuse.Mount(
		mountPoint,
		fuse.FSName("GoFuseFS"),
		fuse.Subtype("fs"),
		fuse.ReadOnly(),
	)
	if err != nil {
		log.Fatalf("Failed to mount: %v", err)
	}
	defer c.Close()
	fmt.Println("Filesystem mounted successfully")

	filesys := &cerebriumfuse.RootFS{Path: nfsDir}
	if err := fs.Serve(c, filesys); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
func addConsoleVibesOnStart() {

	// Reading the content of the file "data.txt"
	data, err := ioutil.ReadFile("./console_art.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Printing the content to the console
	fmt.Println(string(data))
}

func unmountFS(mountPoint string) {
	// Unmount the filesystem
	err := fuse.Unmount(mountPoint)
	if err != nil {
		fmt.Println("Error unmounting filesystem:", err)
	} else {
		fmt.Println("Filesystem unmounted successfully")

	}

}

func clearCacheDirectory(cacheDirPath string) {
	// Get all files in the directory
	files, err := ioutil.ReadDir(cacheDirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Remove each file
	for _, file := range files {
		err := os.Remove(cacheDirPath + "/" + file.Name())
		if err != nil {
			fmt.Println("Error removing file:", err)
		}
	}

	fmt.Println("Files in", cacheDirPath, "cleared successfully")
}
