package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type EC2SpotFleet struct {
	Type       string                 `yaml:"Type"`
	Properties EC2SpotFleetProperties `yaml:"Properties"`
	Condition  interface{}            `yaml:"Condition,omitempty"`
	Metadata   interface{}            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}            `yaml:"DependsOn,omitempty"`
}

type EC2SpotFleetProperties struct {
	SpotFleetRequestConfigData *properties.SpotFleet_SpotFleetRequestConfigData `yaml:"SpotFleetRequestConfigData"`
}

func NewEC2SpotFleet(properties EC2SpotFleetProperties, deps ...interface{}) EC2SpotFleet {
	return EC2SpotFleet{
		Type:       "AWS::EC2::SpotFleet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SpotFleet(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2SpotFleet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SpotFleet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource EC2SpotFleet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SpotFleetProperties) Validate() []error {
	errs := []error{}
	if resource.SpotFleetRequestConfigData == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SpotFleetRequestConfigData'"))
	} else {
		errs = append(errs, resource.SpotFleetRequestConfigData.Validate()...)
	}
	return errs
}
