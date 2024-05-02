package filesystem

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

var (
	CacheDir   string
	cacheSize  = 2                           // Maximum number of entries in the cache
	fileCache  = make(map[string]CacheEntry) // Map from hash to CacheEntry
	cacheMutex sync.RWMutex
)

type CacheEntry struct {
	FilePath string
	LastRead time.Time
}

func HashFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()

	buf := make([]byte, 1024) // Sample size for reading file chunks
	_, err = file.Read(buf)   // Read the first chunk
	if err != nil && err != io.EOF {
		return "", err
	}

	hasher.Write(buf)

	// Seek to the middle of the file
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	if fileSize > 2048 {
		midOffset := fileSize / 2
		file.Seek(midOffset, io.SeekStart)
		file.Read(buf) // Read the middle chunk
		hasher.Write(buf)

		// Seek to the end
		file.Seek(-1024, io.SeekEnd)
		file.Read(buf) // Read the last chunk
		hasher.Write(buf)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func AddFileToCache(filePath string, fileHash string) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// If cache is full, remove the least recently used entry
	if len(fileCache) >= cacheSize {
		var oldestKey string
		var oldestTime = time.Now()

		for k, entry := range fileCache {
			if entry.LastRead.Before(oldestTime) {
				oldestTime = entry.LastRead
				oldestKey = k
			}
		}

		err := fileutils.DeleteFile(fileCache[oldestKey].FilePath)
		if err != nil {
			delete(fileCache, oldestKey)
		} else {
			// @TODO would add some logging here. If this fails the program doesn't need to do anything but we would need to know about it.
		}
	}

	// Do the actual move
	err := fileutils.CopyFile(filePath, path.Join(CacheDir, fileHash))
	if err == nil {
		// Add the new entry
		fileCache[fileHash] = CacheEntry{
			FilePath: path.Join(CacheDir, fileHash),
			LastRead: time.Now(),
		}
	} else {
		// @TODO would add some logging here. If this fails the program doesn't need to do anything but we would need to know about it.
	}

	return nil
}

func GetFileFromCache(fileHash string) (string, bool) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	entry, exists := fileCache[fileHash]
	if exists {
		// Update the timestamp
		fileCache[fileHash] = CacheEntry{
			FilePath: entry.FilePath,
			LastRead: time.Now(),
		}
		return entry.FilePath, true
	}

	return "", false
}
