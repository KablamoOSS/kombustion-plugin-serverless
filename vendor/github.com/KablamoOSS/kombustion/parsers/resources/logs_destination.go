package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type LogsDestination struct {
	Type       string                    `yaml:"Type"`
	Properties LogsDestinationProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

type LogsDestinationProperties struct {
	DestinationName   interface{} `yaml:"DestinationName"`
	DestinationPolicy interface{} `yaml:"DestinationPolicy"`
	RoleArn           interface{} `yaml:"RoleArn"`
	TargetArn         interface{} `yaml:"TargetArn"`
}

func NewLogsDestination(properties LogsDestinationProperties, deps ...interface{}) LogsDestination {
	return LogsDestination{
		Type:       "AWS::Logs::Destination",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLogsDestination(name string, data string) (cf types.TemplateObject, err error) {
	var resource LogsDestination
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LogsDestination - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource LogsDestination) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LogsDestinationProperties) Validate() []error {
	errs := []error{}
	if resource.DestinationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DestinationName'"))
	}
	if resource.DestinationPolicy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DestinationPolicy'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.TargetArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetArn'"))
	}
	return errs
}
