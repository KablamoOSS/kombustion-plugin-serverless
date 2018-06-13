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

// GlueJob Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-job.html
type GlueJob struct {
	Type       string            `yaml:"Type"`
	Properties GlueJobProperties `yaml:"Properties"`
	Condition  interface{}       `yaml:"Condition,omitempty"`
	Metadata   interface{}       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}       `yaml:"DependsOn,omitempty"`
}

// GlueJob Properties
type GlueJobProperties struct {
	AllocatedCapacity interface{}                      `yaml:"AllocatedCapacity,omitempty"`
	DefaultArguments  interface{}                      `yaml:"DefaultArguments,omitempty"`
	Description       interface{}                      `yaml:"Description,omitempty"`
	LogUri            interface{}                      `yaml:"LogUri,omitempty"`
	MaxRetries        interface{}                      `yaml:"MaxRetries,omitempty"`
	Name              interface{}                      `yaml:"Name,omitempty"`
	Role              interface{}                      `yaml:"Role"`
	Command           *properties.JobJobCommand        `yaml:"Command"`
	ExecutionProperty *properties.JobExecutionProperty `yaml:"ExecutionProperty,omitempty"`
	Connections       *properties.JobConnectionsList   `yaml:"Connections,omitempty"`
}

// NewGlueJob constructor creates a new GlueJob
func NewGlueJob(properties GlueJobProperties, deps ...interface{}) GlueJob {
	return GlueJob{
		Type:       "AWS::Glue::Job",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueJob parses GlueJob
func ParseGlueJob(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource GlueJob
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueJob - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseGlueJob validator
func (resource GlueJob) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueJobProperties validator
func (resource GlueJobProperties) Validate() []error {
	errs := []error{}
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Command == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Command'"))
	} else {
		errs = append(errs, resource.Command.Validate()...)
	}
	return errs
}
