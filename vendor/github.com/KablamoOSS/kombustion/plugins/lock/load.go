package lock

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/KablamoOSS/yaml"
)

// FindAndLoadLock - Search the current directory for a Lock file, and load it
func FindAndLoadLock() (Lock Lock, err error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return findAndLoadLock(path)
}

// findAndLoadLock - Search the given directory for a Lock , and load it
// This is seperated to allow for easy testing
func findAndLoadLock(path string) (Lock Lock, err error) {
	var LockPath string
	foundLock := false

	if _, err := os.Stat(path + "/kombustion.lock"); err == nil {
		LockPath = path + "/kombustion.lock"
		foundLock = true
	}

	if foundLock {
		// Read the Lock file
		data, err := ioutil.ReadFile(LockPath)
		if err != nil {
			return Lock, err
		}

		Lock, err := loadLockFromString(data)
		if err != nil {
			return Lock, err
		}
		return Lock, err
	}
	// We didn't find a lock file
	return Lock, nil
}

// loadLockFromString - Load a Lock from a string into a Lock struct
func loadLockFromString(LockYaml []byte) (Lock, error) {
	Lock := Lock{}

	err := yaml.Unmarshal([]byte(LockYaml), &Lock)
	if err != nil {
		return Lock, err
	}
	return Lock, nil
}
