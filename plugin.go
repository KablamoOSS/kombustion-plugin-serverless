package main

import (
	"github.com/KablamoOSS/kombustion-plugin-serverless/resources"
	"github.com/KablamoOSS/kombustion/plugins"
)

var (
	version string
)

// Plugin - Kombustion Serverless Plugin
func Plugin() (plugin plugins.KombustionPlugin, err error) {
	plugin = plugins.KombustionPlugin{
		Version: version,
		Resources: plugins.ParserFunctions{
			"Kablamo::Serverless::Function::Permission": resources.ParseLambdaPermission,
		},
		Outputs:  plugins.ParserFunctions{},
		Mappings: plugins.ParserFunctions{},
		Help: plugins.PluginHelp{
			Description: "A Serverless Plugin",
			TypeMappings: []plugins.TypeMapping{
				{
					Name:        "Kablamo::Serverless::Function::Permission",
					Description: "Creates a permission for a function.",
					Config:      resources.LambdaPermissionConfig{},
				},
			},
		},
	}

	return
}

func main() {}
