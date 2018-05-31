package main

import (
	"github.com/KablamoOSS/kombustion-plugin-serverless/resources"
	"github.com/KablamoOSS/kombustion/plugins"
)

var (
	version string
)

// Plugin - Kombustion Serverless Plugin
func Plugin(plugin *plugins.KombustionPlugin) (interface{}, error) {
	plugin = plugins.KombustionPlugin{
		Version: func() string { return version },
		Resources: func() plugins.ParserFunctions {
			return plugins.ParserFunctions{
				"Kablamo::Serverless::Function::Permission": resources.ParseLambdaPermission,
			}
		},
		Outputs: func() plugins.ParserFunctions {
			return plugins.ParserFunctions{}
		},
		Mappings: func() plugins.ParserFunctions {
			return plugins.ParserFunctions{}
		},
		Help: func() plugins.PluginHelp {
			return plugins.PluginHelp{
				Description: "A Serverless Plugin",
				TypeMappings: []plugins.TypeMapping{
					{
						Name:        "Kablamo::Serverless::Function::Permission",
						Description: "Creates a permission for a function.",
						Config:      resources.LambdaPermissionConfig{},
					},
				},
			}
		},
	}

	return plugin, nil
}

func main() {}
