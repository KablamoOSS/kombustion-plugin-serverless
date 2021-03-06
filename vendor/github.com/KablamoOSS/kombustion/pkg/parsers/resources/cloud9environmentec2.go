package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// Cloud9EnvironmentEC2 Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloud9-environmentec2.html
type Cloud9EnvironmentEC2 struct {
	Type       string                         `yaml:"Type"`
	Properties Cloud9EnvironmentEC2Properties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// Cloud9EnvironmentEC2 Properties
type Cloud9EnvironmentEC2Properties struct {
	AutomaticStopTimeMinutes interface{} `yaml:"AutomaticStopTimeMinutes,omitempty"`
	Description              interface{} `yaml:"Description,omitempty"`
	InstanceType             interface{} `yaml:"InstanceType"`
	Name                     interface{} `yaml:"Name,omitempty"`
	OwnerArn                 interface{} `yaml:"OwnerArn,omitempty"`
	SubnetId                 interface{} `yaml:"SubnetId,omitempty"`
	Repositories             interface{} `yaml:"Repositories,omitempty"`
}

// NewCloud9EnvironmentEC2 constructor creates a new Cloud9EnvironmentEC2
func NewCloud9EnvironmentEC2(properties Cloud9EnvironmentEC2Properties, deps ...interface{}) Cloud9EnvironmentEC2 {
	return Cloud9EnvironmentEC2{
		Type:       "AWS::Cloud9::EnvironmentEC2",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloud9EnvironmentEC2 parses Cloud9EnvironmentEC2
func ParseCloud9EnvironmentEC2(name string, data string) (cf types.TemplateObject, err error) {
	var resource Cloud9EnvironmentEC2
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Cloud9EnvironmentEC2 - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCloud9EnvironmentEC2 validator
func (resource Cloud9EnvironmentEC2) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloud9EnvironmentEC2Properties validator
func (resource Cloud9EnvironmentEC2Properties) Validate() []error {
	errs := []error{}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
