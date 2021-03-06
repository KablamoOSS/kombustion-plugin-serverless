package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2VPCDHCPOptionsAssociation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
type EC2VPCDHCPOptionsAssociation struct {
	Type       string                                 `yaml:"Type"`
	Properties EC2VPCDHCPOptionsAssociationProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// EC2VPCDHCPOptionsAssociation Properties
type EC2VPCDHCPOptionsAssociationProperties struct {
	DhcpOptionsId interface{} `yaml:"DhcpOptionsId"`
	VpcId         interface{} `yaml:"VpcId"`
}

// NewEC2VPCDHCPOptionsAssociation constructor creates a new EC2VPCDHCPOptionsAssociation
func NewEC2VPCDHCPOptionsAssociation(properties EC2VPCDHCPOptionsAssociationProperties, deps ...interface{}) EC2VPCDHCPOptionsAssociation {
	return EC2VPCDHCPOptionsAssociation{
		Type:       "AWS::EC2::VPCDHCPOptionsAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPCDHCPOptionsAssociation parses EC2VPCDHCPOptionsAssociation
func ParseEC2VPCDHCPOptionsAssociation(name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPCDHCPOptionsAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPCDHCPOptionsAssociation - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2VPCDHCPOptionsAssociation validator
func (resource EC2VPCDHCPOptionsAssociation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPCDHCPOptionsAssociationProperties validator
func (resource EC2VPCDHCPOptionsAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.DhcpOptionsId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DhcpOptionsId'"))
	}
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	return errs
}
