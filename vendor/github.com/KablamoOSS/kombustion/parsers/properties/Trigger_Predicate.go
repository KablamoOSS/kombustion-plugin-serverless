package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Trigger_Predicate struct {
	Logical    interface{} `yaml:"Logical,omitempty"`
	Conditions interface{} `yaml:"Conditions,omitempty"`
}

func (resource Trigger_Predicate) Validate() []error {
	errs := []error{}

	return errs
}
