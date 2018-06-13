package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2VPNGatewayRoutePropagation Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type EC2VPNGatewayRoutePropagation struct {
	Type       string                                  `yaml:"Type"`
	Properties EC2VPNGatewayRoutePropagationProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

// EC2VPNGatewayRoutePropagation Properties
type EC2VPNGatewayRoutePropagationProperties struct {
	VpnGatewayId  interface{} `yaml:"VpnGatewayId"`
	RouteTableIds interface{} `yaml:"RouteTableIds"`
}

// NewEC2VPNGatewayRoutePropagation constructor creates a new EC2VPNGatewayRoutePropagation
func NewEC2VPNGatewayRoutePropagation(properties EC2VPNGatewayRoutePropagationProperties, deps ...interface{}) EC2VPNGatewayRoutePropagation {
	return EC2VPNGatewayRoutePropagation{
		Type:       "AWS::EC2::VPNGatewayRoutePropagation",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VPNGatewayRoutePropagation parses EC2VPNGatewayRoutePropagation
func ParseEC2VPNGatewayRoutePropagation(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource EC2VPNGatewayRoutePropagation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VPNGatewayRoutePropagation - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEC2VPNGatewayRoutePropagation validator
func (resource EC2VPNGatewayRoutePropagation) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VPNGatewayRoutePropagationProperties validator
func (resource EC2VPNGatewayRoutePropagationProperties) Validate() []error {
	errs := []error{}
	if resource.VpnGatewayId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpnGatewayId'"))
	}
	if resource.RouteTableIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RouteTableIds'"))
	}
	return errs
}
