package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
	"log"
)

type DAXParameterGroup struct {
	Type       string                      `yaml:"Type"`
	Properties DAXParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DAXParameterGroupProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	ParameterGroupName interface{} `yaml:"ParameterGroupName,omitempty"`
	ParameterNameValues interface{} `yaml:"ParameterNameValues,omitempty"`
}

func NewDAXParameterGroup(properties DAXParameterGroupProperties, deps ...interface{}) DAXParameterGroup {
	return DAXParameterGroup{
		Type:       "AWS::DAX::ParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDAXParameterGroup(name string, data string) (cf plugins.ValueMap, err error) {
	var resource DAXParameterGroup
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DAXParameterGroup - ", err)
		}
		return
	}
	cf = plugins.ValueMap{name: resource}
	return
}

func (resource DAXParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DAXParameterGroupProperties) Validate() []error {
	errs := []error{}
	return errs
}
