package lock

import manifestType "github.com/KablamoOSS/kombustion/manifest"

// UpdateLock - update and write out a new lock file
func UpdateLock(manifest manifestType.Manifest, lockFile Lock) error {
	// First load the lock file
	lockFile, err := FindAndLoadLock()
	if err != nil {
		// Handle if there is no lock file
		return err
	}

	err = updateLock(manifest, lockFile)

	return err
}

func updateLock(manifest manifestType.Manifest, lockFile Lock) error {
	// TODO: reconcile the manifest with the lock file

	err := WriteLockToDisk(lockFile)
	return err
}
