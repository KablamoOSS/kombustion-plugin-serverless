package main

import (
	"github.com/KablamoOSS/kombustion-plugin-serverless/resources"
	"github.com/KablamoOSS/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
	"Kablamo::Serverless::Function::Permission": resources.ParseLambdaPermission,
}

var Outputs = map[string]types.ParserFunc{}

var Mappings = map[string]types.ParserFunc{}

var Help = types.PluginHelp{
	Description: "A Serverless Plugin",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Kablamo::Serverless::Function::Permission",
			Description: "Creates a permission for a function.",
			Config:      resources.LambdaPermissionConfig{},
		},
	},
}

func main() {}
