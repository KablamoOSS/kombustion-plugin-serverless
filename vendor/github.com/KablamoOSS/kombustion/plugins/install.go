package plugins

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/KablamoOSS/kombustion/plugins/lock"
	"github.com/KablamoOSS/kombustion/printer"
	"github.com/mholt/archiver"
	log "github.com/sirupsen/logrus"
)

// installPlugins - Get the lock file and then call installPluginsWithLock
func installPlugins() error {

	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		return err
	}
	lockFile, installErrors := installPluginsWithLock(lockFile)
	if len(installErrors) > 0 {
		for _, err := range installErrors {
			log.Error(err)
		}
	}
	err = lock.WriteLockToDisk(lockFile)
	if err != nil {
		return err
	}
	return nil
}

// installPluginsWithLock - Using a provided lockFile install plugins
func installPluginsWithLock(lockFile lock.Lock) (updatedLockFile lock.Lock, installErrors []error) {
	printer.Step("Installing plugins")
	updatedLockFile = lockFile

	ensureLocalPluginDir()

	for i, plugin := range lockFile.Plugins {
		printer.SubStep(fmt.Sprintf("Installing %s", plugin.Name), 1, false)
		updatedPlugin, errs := installPlugin(plugin)
		if errs != nil {
			installErrors = append(installErrors, errs...)
		}
		updatedLockFile.Plugins[i] = updatedPlugin
	}

	if len(installErrors) > 0 {
		printer.Error("Failed to install plugins")
		return updatedLockFile, installErrors
	}
	printer.Finish("Installed plugins")
	return updatedLockFile, installErrors
}

// installPlugin - Install an individual plugin
func installPlugin(plugin lock.Plugin) (updatedPlugin lock.Plugin, installErrors []error) {
	updatedPlugin = plugin
	for i, resolved := range plugin.Resolved {
		pluginIsInstalled := checkPluginIsAlreadyInstalled(plugin, resolved)
		if pluginIsInstalled {
			printer.SubStep(fmt.Sprintf(
				"Already installed %s@%s[%s/%s]",
				plugin.Name,
				plugin.Version,
				resolved.OperatingSystem,
				resolved.Architecture,
			), 2, false)
		} else {
			var couldInstallFromCache bool
			updatedResolved := resolved

			// Check the local cache for a file
			foundInCache, cacheFile, err := findPluginInCache(plugin, resolved)
			if err != nil {
				installErrors = append(installErrors, err)
			}

			if foundInCache {
				printer.SubStep(fmt.Sprintf(
					"Found cache for %s@%s[%s/%s]",
					plugin.Name,
					plugin.Version,
					resolved.OperatingSystem,
					resolved.Architecture,
				), 2, false)

				var cacheErrors []error
				couldInstallFromCache, cacheErrors = installFromCache(cacheFile, plugin, resolved)
				if len(cacheErrors) != 0 {
					installErrors = append(installErrors, cacheErrors...)
				}
			}
			if couldInstallFromCache == false {
				var downloadErrors []error
				updatedResolved, downloadErrors = downloadPlugin(plugin, resolved)
				if len(downloadErrors) != 0 {
					installErrors = append(installErrors, downloadErrors...)
				}
			}
			updatedPlugin.Resolved[i] = updatedResolved

		}
	}

	printer.SubStep(fmt.Sprintf("Installed %s", plugin.Name), 2, true)
	return updatedPlugin, installErrors
}

func checkPluginIsAlreadyInstalled(plugin lock.Plugin, resolved lock.PluginResolution) (pluginIsInstalled bool) {
	hash, _ := getHashOfFile(resolved.PathOnDisk)
	if hash == resolved.Hash {
		pluginIsInstalled = true
	}
	return pluginIsInstalled
}

// downloadPlugin - Download a single plugin
func downloadPlugin(plugin lock.Plugin, resolved lock.PluginResolution) (updatedResolved lock.PluginResolution, downloadErrors []error) {
	printer.Print(
		fmt.Sprintf(
			"Downloading %s@%s[%s/%s]",
			plugin.Name,
			plugin.Version,
			resolved.OperatingSystem,
			resolved.Architecture,
		),
	)

	// TODO: check the URL is valid

	urlSplit := strings.Split(resolved.URL, "/")
	fileName := urlSplit[len(urlSplit)-1]
	filePath := fmt.Sprintf("%s/%s", getDownloadDir(plugin.Name, plugin.Version), fileName)

	resolved.ArchiveName = fileName

	output, err := os.Create(filePath)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	defer output.Close()

	response, err := http.Get(resolved.URL)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	defer response.Body.Close()

	// Save the downloaded file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}

	resolved.ArchiveHash, err = getHashOfFile(filePath)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	extractedFileName, extractErr := extractPlugin(
		plugin.Name,
		resolved.OperatingSystem,
		resolved.Architecture,
		plugin.Version,
		filePath,
	)
	if extractErr != nil {
		downloadErrors = append(downloadErrors, extractErr)
	}

	resolved.PathOnDisk = extractedFileName

	resolved.Hash, err = getHashOfFile(resolved.PathOnDisk)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}

	printer.SubStep(
		fmt.Sprintf(
			"Downloaded %s@%s[%s/%s]",
			plugin.Name,
			plugin.Version,
			resolved.OperatingSystem,
			resolved.Architecture,
		),
		2, false)

	updatedResolved = resolved
	return updatedResolved, downloadErrors
}

// Extract a plugin downloaded to the local cache, into the local plugin dir
func extractPlugin(pluginName string, operatingSystem string, architecture string, version string, fileName string) (extractedFilePath string, err error) {
	destination := getLocalPluginDir(pluginName, operatingSystem, architecture, version)
	extracter := getExtracter(fileName)
	if extracter == nil {
		return extractedFilePath, fmt.Errorf(fmt.Sprintf("Unable to extract: %s", fileName))
	}
	err = extracter.Open(fileName, destination)
	var found bool
	var extractedFileName string
	extractedFileName, found, err = findExtractedPluginFile(destination)
	if found == false {
		return extractedFilePath, fmt.Errorf("Invalid plugin archive: %s", pluginName)
	}
	extractedFilePath = fmt.Sprintf("%s/%s", destination, extractedFileName)
	return extractedFilePath, err
}

func findExtractedPluginFile(directory string) (fileName string, found bool, err error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return fileName, found, err
	}

	for _, file := range files {
		fileName = file.Name()
		if strings.HasSuffix(fileName, ".so") {
			found = true
			return fileName, found, err
		}
	}
	return fileName, found, err
}

// getPluginDir - Get the plugin directory for this project,
// and make it if it doesn't exist
func getLocalPluginDir(pluginName string, operatingSystem string, architecture string, version string) string {
	pluginDir := fmt.Sprintf(".kombustion/plugins/%s/%s/%s/%s", pluginName, version, operatingSystem, architecture)
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}

func ensureLocalPluginDir() string {
	pluginDir := fmt.Sprintf(".kombustion/plugins/")
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}

// Get the download cache directory
func getDownloadDir(pluginName string, version string) string {
	cacheDir := GetCacheDir()
	pluginDir := fmt.Sprintf("%s/%s/%s", cacheDir, pluginName, version)
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}

// Clean up the download directory
func cleanDownloadDir() error {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	downloadDir := fmt.Sprintf("%s/.kombustion/cache", usr.HomeDir)

	return os.RemoveAll(downloadDir)
}

// The PR hasn't been released into a new version, so we're putting it here
// https://github.com/mholt/archiver/pull/45/files
func getExtracter(fpath string) archiver.Archiver {
	for _, fmt := range archiver.SupportedFormats {
		if fmt.Match(fpath) {
			return fmt
		}
	}
	return nil
}
