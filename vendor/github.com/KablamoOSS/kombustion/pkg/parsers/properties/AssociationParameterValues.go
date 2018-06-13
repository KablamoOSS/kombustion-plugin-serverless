package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// AssociationParameterValues Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-association-parametervalues.html
type AssociationParameterValues struct {
	ParameterValues interface{} `yaml:"ParameterValues"`
}

// AssociationParameterValues validation
func (resource AssociationParameterValues) Validate() []error {
	errs := []error{}

	if resource.ParameterValues == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterValues'"))
	}
	return errs
}
