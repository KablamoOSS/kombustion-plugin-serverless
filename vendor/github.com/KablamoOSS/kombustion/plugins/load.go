package plugins

import (
	"plugin"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/plugins/lock"
)

// LoadPlugins -
func LoadPlugins() (resources, outputs, mappings map[string]ParserFunc) {
	resources, outputs, mappings = make(map[string]ParserFunc), make(map[string]ParserFunc), make(map[string]ParserFunc)

	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		log.Fatal(err)
	}

	for _, plugin := range lockFile.Plugins {

		for _, resolved := range plugin.Resolved {
			if runtime.GOOS == resolved.OperatingSystem &&
				runtime.GOARCH == resolved.Architecture {
				loadPlugin(resolved.PathOnDisk, resources, outputs, mappings)
			}
		}
	}

	return resources, outputs, mappings
}

func loadPlugin(fileName string, resources, outputs, mappings map[string]ParserFunc) {
	p, err := plugin.Open(fileName)

	if err != nil {
		log.WithFields(log.Fields{
			"filename": fileName,
			"err":      err,
		}).Warn("error reading plugin file")
		return
	}

	// Load the plugins Resources
	pluginResources, err := p.Lookup("Resources")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": fileName,
			"err":      err,
		}).Warn("error reading resource plugin")
	}
	for k, v := range *pluginResources.(*map[string]ParserFunc) {
		if _, ok := resources[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"resource": k,
			}).Warn("duplicate resource definition for resource")
		} else {
			resources[k] = v
		}
	}

	// Load the plugins Outputs
	pluginOutputs, err := p.Lookup("Outputs")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": fileName,
			"err":      err,
		}).Warn("error reading resource plugin")
	}
	for k, v := range *pluginOutputs.(*map[string]ParserFunc) {
		if _, ok := outputs[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"output": k,
			}).Warn("duplicate output definition for output")
		} else {
			outputs[k] = v
		}
	}

	// Load the plugins Mappings
	pluginMappings, err := p.Lookup("Mappings")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": fileName,
			"err":      err,
		}).Warn("error reading resource plugin")
	}
	for k, v := range *pluginMappings.(*map[string]ParserFunc) {
		if _, ok := mappings[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"mapping": k,
			}).Warn("duplicate mapping definition for mapping")
		} else {
			mappings[k] = v
		}
	}
	return
}
