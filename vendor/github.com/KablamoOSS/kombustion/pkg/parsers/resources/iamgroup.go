package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// IAMGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
type IAMGroup struct {
	Type       string             `yaml:"Type"`
	Properties IAMGroupProperties `yaml:"Properties"`
	Condition  interface{}        `yaml:"Condition,omitempty"`
	Metadata   interface{}        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

// IAMGroup Properties
type IAMGroupProperties struct {
	GroupName         interface{} `yaml:"GroupName,omitempty"`
	Path              interface{} `yaml:"Path,omitempty"`
	ManagedPolicyArns interface{} `yaml:"ManagedPolicyArns,omitempty"`
	Policies          interface{} `yaml:"Policies,omitempty"`
}

// NewIAMGroup constructor creates a new IAMGroup
func NewIAMGroup(properties IAMGroupProperties, deps ...interface{}) IAMGroup {
	return IAMGroup{
		Type:       "AWS::IAM::Group",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseIAMGroup parses IAMGroup
func ParseIAMGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource IAMGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IAMGroup - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseIAMGroup validator
func (resource IAMGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseIAMGroupProperties validator
func (resource IAMGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
