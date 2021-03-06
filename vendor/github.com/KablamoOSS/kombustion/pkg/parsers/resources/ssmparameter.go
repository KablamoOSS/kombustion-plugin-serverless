package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SSMParameter Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html
type SSMParameter struct {
	Type       string                 `yaml:"Type"`
	Properties SSMParameterProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

// SSMParameter Properties
type SSMParameterProperties struct {
	AllowedPattern interface{} `yaml:"AllowedPattern,omitempty"`
	Description    interface{} `yaml:"Description,omitempty"`
	Name           interface{} `yaml:"Name,omitempty"`
	Type           interface{} `yaml:"Type"`
	Value          interface{} `yaml:"Value"`
}

// NewSSMParameter constructor creates a new SSMParameter
func NewSSMParameter(properties SSMParameterProperties, deps ...interface{}) SSMParameter {
	return SSMParameter{
		Type:       "AWS::SSM::Parameter",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMParameter parses SSMParameter
func ParseSSMParameter(name string, data string) (cf types.TemplateObject, err error) {
	var resource SSMParameter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMParameter - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSSMParameter validator
func (resource SSMParameter) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMParameterProperties validator
func (resource SSMParameterProperties) Validate() []error {
	errs := []error{}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
