package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ElasticLoadBalancingV2ListenerRule Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html
type ElasticLoadBalancingV2ListenerRule struct {
	Type       string                                       `yaml:"Type"`
	Properties ElasticLoadBalancingV2ListenerRuleProperties `yaml:"Properties"`
	Condition  interface{}                                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                  `yaml:"DependsOn,omitempty"`
}

// ElasticLoadBalancingV2ListenerRule Properties
type ElasticLoadBalancingV2ListenerRuleProperties struct {
	ListenerArn interface{} `yaml:"ListenerArn"`
	Priority    interface{} `yaml:"Priority"`
	Actions     interface{} `yaml:"Actions"`
	Conditions  interface{} `yaml:"Conditions"`
}

// NewElasticLoadBalancingV2ListenerRule constructor creates a new ElasticLoadBalancingV2ListenerRule
func NewElasticLoadBalancingV2ListenerRule(properties ElasticLoadBalancingV2ListenerRuleProperties, deps ...interface{}) ElasticLoadBalancingV2ListenerRule {
	return ElasticLoadBalancingV2ListenerRule{
		Type:       "AWS::ElasticLoadBalancingV2::ListenerRule",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElasticLoadBalancingV2ListenerRule parses ElasticLoadBalancingV2ListenerRule
func ParseElasticLoadBalancingV2ListenerRule(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ElasticLoadBalancingV2ListenerRule
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticLoadBalancingV2ListenerRule - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseElasticLoadBalancingV2ListenerRule validator
func (resource ElasticLoadBalancingV2ListenerRule) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElasticLoadBalancingV2ListenerRuleProperties validator
func (resource ElasticLoadBalancingV2ListenerRuleProperties) Validate() []error {
	errs := []error{}
	if resource.ListenerArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ListenerArn'"))
	}
	if resource.Priority == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Priority'"))
	}
	if resource.Actions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Actions'"))
	}
	if resource.Conditions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Conditions'"))
	}
	return errs
}