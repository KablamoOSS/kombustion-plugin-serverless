package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// WAFRegionalSqlInjectionMatchSet Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html
type WAFRegionalSqlInjectionMatchSet struct {
	Type       string                                    `yaml:"Type"`
	Properties WAFRegionalSqlInjectionMatchSetProperties `yaml:"Properties"`
	Condition  interface{}                               `yaml:"Condition,omitempty"`
	Metadata   interface{}                               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                               `yaml:"DependsOn,omitempty"`
}

// WAFRegionalSqlInjectionMatchSet Properties
type WAFRegionalSqlInjectionMatchSetProperties struct {
	Name                    interface{} `yaml:"Name"`
	SqlInjectionMatchTuples interface{} `yaml:"SqlInjectionMatchTuples,omitempty"`
}

// NewWAFRegionalSqlInjectionMatchSet constructor creates a new WAFRegionalSqlInjectionMatchSet
func NewWAFRegionalSqlInjectionMatchSet(properties WAFRegionalSqlInjectionMatchSetProperties, deps ...interface{}) WAFRegionalSqlInjectionMatchSet {
	return WAFRegionalSqlInjectionMatchSet{
		Type:       "AWS::WAFRegional::SqlInjectionMatchSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWAFRegionalSqlInjectionMatchSet parses WAFRegionalSqlInjectionMatchSet
func ParseWAFRegionalSqlInjectionMatchSet(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalSqlInjectionMatchSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalSqlInjectionMatchSet - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseWAFRegionalSqlInjectionMatchSet validator
func (resource WAFRegionalSqlInjectionMatchSet) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWAFRegionalSqlInjectionMatchSetProperties validator
func (resource WAFRegionalSqlInjectionMatchSetProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
