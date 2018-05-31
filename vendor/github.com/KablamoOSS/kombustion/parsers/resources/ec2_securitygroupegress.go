package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
	"log"
	"fmt"
)

type EC2SecurityGroupEgress struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SecurityGroupEgressProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SecurityGroupEgressProperties struct {
	CidrIp interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6 interface{} `yaml:"CidrIpv6,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	DestinationPrefixListId interface{} `yaml:"DestinationPrefixListId,omitempty"`
	DestinationSecurityGroupId interface{} `yaml:"DestinationSecurityGroupId,omitempty"`
	FromPort interface{} `yaml:"FromPort,omitempty"`
	GroupId interface{} `yaml:"GroupId"`
	IpProtocol interface{} `yaml:"IpProtocol"`
	ToPort interface{} `yaml:"ToPort,omitempty"`
}

func NewEC2SecurityGroupEgress(properties EC2SecurityGroupEgressProperties, deps ...interface{}) EC2SecurityGroupEgress {
	return EC2SecurityGroupEgress{
		Type:       "AWS::EC2::SecurityGroupEgress",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SecurityGroupEgress(name string, data string) (cf plugins.ValueMap, err error) {
	var resource EC2SecurityGroupEgress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SecurityGroupEgress - ", err)
		}
		return
	}
	cf = plugins.ValueMap{name: resource}
	return
}

func (resource EC2SecurityGroupEgress) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SecurityGroupEgressProperties) Validate() []error {
	errs := []error{}
	if resource.GroupId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupId'"))
	}
	if resource.IpProtocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errs
}
