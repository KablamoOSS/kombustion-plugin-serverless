package plugins

// KombustionPlugin - A plugin struct for Kombustion
// All of the following are required for a valid plugin
type KombustionPlugin interface {
	// The resources/outputs/mapping
	Resources() ParserFunctions
	Outputs() ParserFunctions
	Mappings() ParserFunctions

	// Help information for the user of this plugin
	Help() PluginHelp

	// The version of the plugin
	Version() string
}

// PluginHelp - a set of available documentation fields
type PluginHelp struct {
	// The name of the plugin
	Name string
	// A short description of what the plugin does
	Description string

	// Help information for all the types this pplugin provides
	TypeMappings []TypeMapping

	// Examples/Snippets of how this plugin can be used
	Snippets []string
}

// ValueMap - The return type for ParserFunc
type ValueMap map[string]interface{}

// ParserFunctions - The functions called to process resources, outputs and mappings
type ParserFunctions map[string]ParserFunc

// ParserFunc - a definition of the function called for resource/output/mapping parsers
type ParserFunc func(string, string) (ValueMap, error)

// TypeMapping - recursive list of types with its associated config object
type TypeMapping struct {
	Name        string
	Description string
	Config      interface{}
}
