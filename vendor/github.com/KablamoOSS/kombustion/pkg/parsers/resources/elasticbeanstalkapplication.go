package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ElasticBeanstalkApplication Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type ElasticBeanstalkApplication struct {
	Type       string                                `yaml:"Type"`
	Properties ElasticBeanstalkApplicationProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// ElasticBeanstalkApplication Properties
type ElasticBeanstalkApplicationProperties struct {
	ApplicationName         interface{}                                               `yaml:"ApplicationName,omitempty"`
	Description             interface{}                                               `yaml:"Description,omitempty"`
	ResourceLifecycleConfig *properties.ApplicationApplicationResourceLifecycleConfig `yaml:"ResourceLifecycleConfig,omitempty"`
}

// NewElasticBeanstalkApplication constructor creates a new ElasticBeanstalkApplication
func NewElasticBeanstalkApplication(properties ElasticBeanstalkApplicationProperties, deps ...interface{}) ElasticBeanstalkApplication {
	return ElasticBeanstalkApplication{
		Type:       "AWS::ElasticBeanstalk::Application",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticBeanstalkApplication parses ElasticBeanstalkApplication
func ParseElasticBeanstalkApplication(name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticBeanstalkApplication
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkApplication - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticBeanstalkApplication validator
func (resource ElasticBeanstalkApplication) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticBeanstalkApplicationProperties validator
func (resource ElasticBeanstalkApplicationProperties) Validate() []error {
	errs := []error{}
	return errs
}
