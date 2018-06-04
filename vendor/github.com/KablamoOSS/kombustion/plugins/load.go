package plugins

import (
	"fmt"
	"os"
	"plugin"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/plugins/lock"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
)

// LoadPlugins -
func LoadPlugins() (resources, outputs, mappings map[string]kombustionTypes.ParserFunc) {
	resources, outputs, mappings =
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc)

	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		log.Fatal(err)
	}

	for _, plugin := range lockFile.Plugins {

		for _, resolved := range plugin.Resolved {
			if runtime.GOOS == resolved.OperatingSystem &&
				runtime.GOARCH == resolved.Architecture {
				loadPlugin(plugin.Name, plugin.Version, resolved.PathOnDisk, resources, outputs, mappings)
			}
		}
	}

	return resources, outputs, mappings
}

func loadPlugin(pluginName string, pluginVersion string, pluginPath string, resources, outputs, mappings map[string]kombustionTypes.ParserFunc) {
	resources, outputs, mappings =
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc),
		make(map[string]kombustionTypes.ParserFunc)

	if !pluginExists(pluginPath) {
		fmt.Fprintf(os.Stderr, "error: invalid plugin file: %s\n", pluginPath)
		os.Exit(1)
	}

	// open the plug-in file, causing its package init function to run,
	// thereby registering the module
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: failed to load plugin: %v\n", err)
	}

	// Help
	// helpConstructor, _ := p.Lookup("Help")
	// helpBlob := *helpConstructor.(*func() bytes.Buffer)
	// help, _ := LoadHelp(helpBlob())

	// fmt.Printf("help: %+v \n\n", help)

	// Resources
	resourcesConstructor, _ := p.Lookup("Resources")
	resourcesFuncs := *resourcesConstructor.(*map[string]func(
		ctx map[string]interface{},
		name string,
		data string,
	) []byte)
	fmt.Printf("\n resourcesFuncs: %+v\n\n", resourcesFuncs)

	data := `
Type: Kablamo::Serverless::Function::Permission
Properties:
  Name: "kombustionTestAPI"
  Action: "lambda:InvokeFunction"
  FunctionName: "tere"
  SourceArn: "sdfsdfsdf"
  SourceApiGateway: "asdad"
`
	type AWSKeys struct {
		Key int
	}
	ctx := make(map[string]interface{})
	ctx["AWSKeys"] = AWSKeys{Key: 3}
	for _, n := range resourcesFuncs {

		r1 := n(ctx, "testName", data)

		fmt.Printf("\n %s r1: %+v\n\n\n", pluginName, r1)

		rrr, _ := loadResource(r1)
		fmt.Printf("\n %s: %+v\n\n\n", pluginName, rrr)
		fmt.Print(rrr["testName"])
	}

	// math, err := p.Lookup("Math")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": pluginPath,
	// 		"err":      err,
	// 	}).Warn("error reading resource plugin")
	// }

	// answer := math.(func(a int, y int) int)
	// fmt.Printf("%+v \n", answer(2, 3))

	// register, err := p.Lookup("Register")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: failed to load plugin: %v\n", err)
	// }

	// modGo := *register.(*func() interface{})

	// m := modGo()
	// fmt.Printf("%+v", m)
	// m.(registry.Plugin).Init(context.Background(), types.PluginConfig{})

	// Instantiate a copy of the module registered by the plug-in.
	// modGo := registry.NewPlugin(pluginName)

	// Initialize mod_go
	// modGo.Init(context.Background(), types.PluginConfig{})

	// p, err := plugin.Open(fileName)

	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": fileName,
	// 		"err":      err,
	// 	}).Warn("error reading plugin file")
	// 	return
	// }

	// open the plug-in file, causing its package init function to run,
	// thereby registering the module
	// _, err := plugin.Open(fileName)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: failed to load plugin: %v\n", err)
	// }
	// pluginRegister, err := p.Lookup("Register")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: failed to load plugin: %v\n", err)
	// }

	// // name, version, pluginActual := pluginRegister.(func() (string, string, interface{}))()
	// config := PluginConfig{}

	// name, version := pluginActual.(KombustionPlugin).Init(context.Background(), config)
	// r := pluginActual.(KombustionPlugin).Resources(context.Background(), config)

	// fmt.Printf("Resources: %+v", r)
	// // Instantiate a copy of the module registered by the plug-in.
	// loadedPlugin := LoadPlugin(name, version)

	// fmt.Printf("loadedPlugin: %+v", loadedPlugin)

	// // Initialize mod_go
	// loadedPlugin.Init(context.Background(), config)

	// resourcesLoad, err := p.Lookup("Resources")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": fileName,
	// 		"err":      err,
	// 	}).Warn("error loading plugin")
	// }

	// outputsLoad, err := p.Lookup("Outputs")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": fileName,
	// 		"err":      err,
	// 	}).Warn("error loading plugin")
	// }
	// mappingsLoad, err := p.Lookup("Mappings")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": fileName,
	// 		"err":      err,
	// 	}).Warn("error loading plugin")
	// }

	// helpLoad, err := p.Lookup("Help")
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"filename": fileName,
	// 		"err":      err,
	// 	}).Warn("error loading plugin")
	// }

	// for k, v := range *resourcesLoad.(*map[string]plugins.ParserFunc) {
	// 	if _, ok := resources[k]; ok { // Check for duplicates
	// 		log.WithFields(log.Fields{
	// 			"resource": k,
	// 		}).Warn("duplicate resource definition for resource")
	// 	} else {
	// 		resources[k] = v
	// 	}
	// }

	// // Load the plugins Outputs
	// for k, v := range *outputsLoad.(*map[string]plugins.ParserFunc) {
	// 	if _, ok := outputs[k]; ok { // Check for duplicates
	// 		log.WithFields(log.Fields{
	// 			"output": k,
	// 		}).Warn("duplicate output definition for output")
	// 	} else {
	// 		outputs[k] = v
	// 	}
	// }

	// // Load the plugins Mappings
	// for k, v := range *mappingsLoad.(*map[string]plugins.ParserFunc) {
	// 	if _, ok := mappings[k]; ok { // Check for duplicates
	// 		log.WithFields(log.Fields{
	// 			"mapping": k,
	// 		}).Warn("duplicate mapping definition for mapping")
	// 	} else {
	// 		mappings[k] = v
	// 	}
	// }
	return
}

func pluginExists(filePath string) bool {
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return true
	}
	return false
}
