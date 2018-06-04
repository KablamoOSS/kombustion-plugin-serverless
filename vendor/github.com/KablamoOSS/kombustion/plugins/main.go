// This file exposes an api to the CLI wrapper for the
// CLI tasks required

package plugins

import (
	"log"

	"github.com/KablamoOSS/kombustion/manifest"
	"github.com/KablamoOSS/kombustion/printer"
	"github.com/urfave/cli"
)

// AddPluginToManifest file and update it
func AddPluginToManifest(c *cli.Context) error {
	printer.Print("Kombusting")
	// Try and load the manifest
	manifest, err := manifest.FindAndLoadManifest()
	if err != nil {
		log.Fatal("No kombustion.yaml manifest. Create one with: kombustion init")
	}

	// Get the plugin to add
	plugins := c.Args()

	// Add them
	manifest, err = addPluginsToManifest(manifest, plugins)
	if err != nil {
		log.Fatal(err)
	}

	// Now install them
	err = installPlugins()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// InstallPlugins to reconcile the manifest to the lock file, and then the lock file to the disk
// Which is to say, ensure the manifest and lock file agree with the state of the plugins
// and then ensure the lock file agrees with the disk on the state of the plugins
func InstallPlugins(c *cli.Context) {
	printer.Print("Kombusting")
	err := installPlugins()
	if err != nil {
		log.Fatal(err)
		printer.Error("Install failed")
	}
}
