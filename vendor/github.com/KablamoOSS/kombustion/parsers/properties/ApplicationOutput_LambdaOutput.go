package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type ApplicationOutput_LambdaOutput struct {
	ResourceARN interface{} `yaml:"ResourceARN"`
	RoleARN     interface{} `yaml:"RoleARN"`
}

func (resource ApplicationOutput_LambdaOutput) Validate() []error {
	errs := []error{}

	if resource.ResourceARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceARN'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	return errs
}
