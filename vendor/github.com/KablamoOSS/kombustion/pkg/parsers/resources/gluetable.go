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

// GlueTable Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-table.html
type GlueTable struct {
	Type       string              `yaml:"Type"`
	Properties GlueTableProperties `yaml:"Properties"`
	Condition  interface{}         `yaml:"Condition,omitempty"`
	Metadata   interface{}         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}         `yaml:"DependsOn,omitempty"`
}

// GlueTable Properties
type GlueTableProperties struct {
	CatalogId    interface{}                 `yaml:"CatalogId"`
	DatabaseName interface{}                 `yaml:"DatabaseName"`
	TableInput   *properties.TableTableInput `yaml:"TableInput"`
}

// NewGlueTable constructor creates a new GlueTable
func NewGlueTable(properties GlueTableProperties, deps ...interface{}) GlueTable {
	return GlueTable{
		Type:       "AWS::Glue::Table",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseGlueTable parses GlueTable
func ParseGlueTable(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource GlueTable
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueTable - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseGlueTable validator
func (resource GlueTable) Validate() []error {
	return resource.Properties.Validate()
}

// ParseGlueTableProperties validator
func (resource GlueTableProperties) Validate() []error {
	errs := []error{}
	if resource.CatalogId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CatalogId'"))
	}
	if resource.DatabaseName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseName'"))
	}
	if resource.TableInput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableInput'"))
	} else {
		errs = append(errs, resource.TableInput.Validate()...)
	}
	return errs
}