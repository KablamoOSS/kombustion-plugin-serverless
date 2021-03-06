package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// LambdaVersion Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html
type LambdaVersion struct {
	Type       string                  `yaml:"Type"`
	Properties LambdaVersionProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// LambdaVersion Properties
type LambdaVersionProperties struct {
	CodeSha256   interface{} `yaml:"CodeSha256,omitempty"`
	Description  interface{} `yaml:"Description,omitempty"`
	FunctionName interface{} `yaml:"FunctionName"`
}

// NewLambdaVersion constructor creates a new LambdaVersion
func NewLambdaVersion(properties LambdaVersionProperties, deps ...interface{}) LambdaVersion {
	return LambdaVersion{
		Type:       "AWS::Lambda::Version",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseLambdaVersion parses LambdaVersion
func ParseLambdaVersion(name string, data string) (cf types.TemplateObject, err error) {
	var resource LambdaVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaVersion - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseLambdaVersion validator
func (resource LambdaVersion) Validate() []error {
	return resource.Properties.Validate()
}

// ParseLambdaVersionProperties validator
func (resource LambdaVersionProperties) Validate() []error {
	errs := []error{}
	if resource.FunctionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionName'"))
	}
	return errs
}
