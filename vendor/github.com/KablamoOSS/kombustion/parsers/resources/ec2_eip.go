package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
	"log"
)

type EC2EIP struct {
	Type       string                      `yaml:"Type"`
	Properties EC2EIPProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2EIPProperties struct {
	Domain interface{} `yaml:"Domain,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
}

func NewEC2EIP(properties EC2EIPProperties, deps ...interface{}) EC2EIP {
	return EC2EIP{
		Type:       "AWS::EC2::EIP",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2EIP(name string, data string) (cf plugins.ValueMap, err error) {
	var resource EC2EIP
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2EIP - ", err)
		}
		return
	}
	cf = plugins.ValueMap{name: resource}
	return
}

func (resource EC2EIP) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2EIPProperties) Validate() []error {
	errs := []error{}
	return errs
}
