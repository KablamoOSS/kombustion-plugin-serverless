package main

import (
	"github.com/KablamoOSS/kombustion-plugin-serverless/resources"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api"
	"github.com/KablamoOSS/kombustion/pkg/plugins/api/types"
)

var (
	version string
	name    string
)

func init() {
	if version == "" {
		version = "BUILT_FROM_SOURCE"
	}
	if name == "" {
		name = "kombustion-plugin-serverless"
	}
}

// Register plugin
func Register() []byte {
	return api.RegisterPlugin(types.Config{
		Name:               name,
		Version:            version,
		Prefix:             "Kablamo",
		RequiresAWSSession: false,
		Help: types.Help{
			Description: "A Serverless Plugin",
			TypeMappings: []types.TypeMapping{
				{
					Name:        "Serverless::Function::Permission",
					Description: "Creates a permission for a function.",
					Config:      resources.LambdaPermissionConfig{},
				},
			},
		},
	})
}

// Resources for this plugin
var Resources = map[string]func(
	name string,
	data string,
) []byte{
	"Serverless::Function::Permission": api.RegisterResource(resources.ParseLambdaPermission),
}

// Outputs for this plugin
var Outputs = map[string]func(
	name string,
	data string,
) []byte{}

// Mappings for this plugin
var Mappings = map[string]func(
	name string,
	data string,
) []byte{}

func main() {}
