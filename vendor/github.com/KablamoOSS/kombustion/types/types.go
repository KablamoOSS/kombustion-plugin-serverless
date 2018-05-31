package types

type Validatable interface {
	Validate() []error
}

// CfResource - A Cloudformation Resource
type CfResource struct {
	Type       string      `yaml:"Type"`
	Properties interface{} `yaml:"Properties"`
	Condition  interface{} `yaml:"Condition,omitempty"`
	Metadata   interface{} `yaml:"Metadata,omitempty"`
	DependsOn  interface{} `yaml:"DependsOn,omitempty"`
}

// ResourceMap - a map of resouces
type ResourceMap map[string]CfResource
