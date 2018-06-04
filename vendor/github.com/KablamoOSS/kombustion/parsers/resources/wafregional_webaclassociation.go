package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type WAFRegionalWebACLAssociation struct {
	Type       string                                 `yaml:"Type"`
	Properties WAFRegionalWebACLAssociationProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

type WAFRegionalWebACLAssociationProperties struct {
	ResourceArn interface{} `yaml:"ResourceArn"`
	WebACLId    interface{} `yaml:"WebACLId"`
}

func NewWAFRegionalWebACLAssociation(properties WAFRegionalWebACLAssociationProperties, deps ...interface{}) WAFRegionalWebACLAssociation {
	return WAFRegionalWebACLAssociation{
		Type:       "AWS::WAFRegional::WebACLAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseWAFRegionalWebACLAssociation(name string, data string) (cf types.TemplateObject, err error) {
	var resource WAFRegionalWebACLAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: WAFRegionalWebACLAssociation - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource WAFRegionalWebACLAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource WAFRegionalWebACLAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.ResourceArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceArn'"))
	}
	if resource.WebACLId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'WebACLId'"))
	}
	return errs
}
