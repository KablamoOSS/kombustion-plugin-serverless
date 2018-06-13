package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ElasticBeanstalkApplicationVersion Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type ElasticBeanstalkApplicationVersion struct {
	Type       string                                       `yaml:"Type"`
	Properties ElasticBeanstalkApplicationVersionProperties `yaml:"Properties"`
	Condition  interface{}                                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                  `yaml:"DependsOn,omitempty"`
}

// ElasticBeanstalkApplicationVersion Properties
type ElasticBeanstalkApplicationVersionProperties struct {
	ApplicationName interface{}                                `yaml:"ApplicationName"`
	Description     interface{}                                `yaml:"Description,omitempty"`
	SourceBundle    *properties.ApplicationVersionSourceBundle `yaml:"SourceBundle"`
}

// NewElasticBeanstalkApplicationVersion constructor creates a new ElasticBeanstalkApplicationVersion
func NewElasticBeanstalkApplicationVersion(properties ElasticBeanstalkApplicationVersionProperties, deps ...interface{}) ElasticBeanstalkApplicationVersion {
	return ElasticBeanstalkApplicationVersion{
		Type:       "AWS::ElasticBeanstalk::ApplicationVersion",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticBeanstalkApplicationVersion parses ElasticBeanstalkApplicationVersion
func ParseElasticBeanstalkApplicationVersion(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticBeanstalkApplicationVersion
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticBeanstalkApplicationVersion - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticBeanstalkApplicationVersion validator
func (resource ElasticBeanstalkApplicationVersion) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticBeanstalkApplicationVersionProperties validator
func (resource ElasticBeanstalkApplicationVersionProperties) Validate() []error {
	errs := []error{}
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.SourceBundle == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceBundle'"))
	} else {
		errs = append(errs, resource.SourceBundle.Validate()...)
	}
	return errs
}
