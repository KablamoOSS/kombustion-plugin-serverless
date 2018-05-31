package lock

import (
	"io/ioutil"

	"github.com/KablamoOSS/yaml"
)

// WriteLockToDisk - Write the Lock to disk
func WriteLockToDisk(Lock Lock) error {

	// Mashall the the struct into yaml
	LockString, err := yaml.Marshal(&Lock)
	if err != nil {
		return err
	}

	// Write the LockString
	err = ioutil.WriteFile("kombustion.lock", LockString, 0644)
	if err != nil {
		return err
	}
	return nil
}
