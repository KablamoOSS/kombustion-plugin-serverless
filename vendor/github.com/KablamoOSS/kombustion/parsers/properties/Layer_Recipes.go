package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Layer_Recipes struct {
	Configure interface{} `yaml:"Configure,omitempty"`
	Deploy    interface{} `yaml:"Deploy,omitempty"`
	Setup     interface{} `yaml:"Setup,omitempty"`
	Shutdown  interface{} `yaml:"Shutdown,omitempty"`
	Undeploy  interface{} `yaml:"Undeploy,omitempty"`
}

func (resource Layer_Recipes) Validate() []error {
	errs := []error{}

	return errs
}
