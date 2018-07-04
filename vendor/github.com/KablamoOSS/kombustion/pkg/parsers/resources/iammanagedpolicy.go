package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// IAMManagedPolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
type IAMManagedPolicy struct {
	Type       string                     `yaml:"Type"`
	Properties IAMManagedPolicyProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// IAMManagedPolicy Properties
type IAMManagedPolicyProperties struct {
	Description       interface{} `yaml:"Description,omitempty"`
	ManagedPolicyName interface{} `yaml:"ManagedPolicyName,omitempty"`
	Path              interface{} `yaml:"Path,omitempty"`
	PolicyDocument    interface{} `yaml:"PolicyDocument"`
	Groups            interface{} `yaml:"Groups,omitempty"`
	Roles             interface{} `yaml:"Roles,omitempty"`
	Users             interface{} `yaml:"Users,omitempty"`
}

// NewIAMManagedPolicy constructor creates a new IAMManagedPolicy
func NewIAMManagedPolicy(properties IAMManagedPolicyProperties, deps ...interface{}) IAMManagedPolicy {
	return IAMManagedPolicy{
		Type:       "AWS::IAM::ManagedPolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIAMManagedPolicy parses IAMManagedPolicy
func ParseIAMManagedPolicy(name string, data string) (cf types.TemplateObject, err error) {
	var resource IAMManagedPolicy
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMManagedPolicy - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseIAMManagedPolicy validator
func (resource IAMManagedPolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIAMManagedPolicyProperties validator
func (resource IAMManagedPolicyProperties) Validate() []error {
	errs := []error{}
	if resource.PolicyDocument == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyDocument'"))
	}
	return errs
}
