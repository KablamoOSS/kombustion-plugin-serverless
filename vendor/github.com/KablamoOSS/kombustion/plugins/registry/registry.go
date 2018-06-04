package registry

import (
	"context"
	"fmt"

	"github.com/KablamoOSS/kombustion/plugins/types"
)

// Plugin is the interface implementated by types that
// register themselves as modular plug-ins.
type Plugin interface {

	// Init initializes the module.
	Init(ctx context.Context, config types.PluginConfig)
}

var pluginRegistry = map[string]func() Plugin{}

// RegisterPlugin registers a new module with its name and function
// that returns a new, uninitialized instance of the module type.
func RegisterPlugin(name string, constructor func() Plugin) {
	fmt.Printf("RegisterPlugin: %+v", name)
	pluginRegistry[name] = constructor
}

// NewPlugin instantiates a new instance of the module type with the
// specified name.
func NewPlugin(name string) Plugin {
	fmt.Printf("pluginRegistry: %+v", pluginRegistry)

	return pluginRegistry[name]()
}

func GetRegistry() map[string]func() Plugin {
	return pluginRegistry
}
