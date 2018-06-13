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

// ApiGatewayDeployment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type ApiGatewayDeployment struct {
	Type       string                         `yaml:"Type"`
	Properties ApiGatewayDeploymentProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// ApiGatewayDeployment Properties
type ApiGatewayDeploymentProperties struct {
	Description      interface{}                            `yaml:"Description,omitempty"`
	RestApiId        interface{}                            `yaml:"RestApiId"`
	StageName        interface{}                            `yaml:"StageName,omitempty"`
	StageDescription *properties.DeploymentStageDescription `yaml:"StageDescription,omitempty"`
}

// NewApiGatewayDeployment constructor creates a new ApiGatewayDeployment
func NewApiGatewayDeployment(properties ApiGatewayDeploymentProperties, deps ...interface{}) ApiGatewayDeployment {
	return ApiGatewayDeployment{
		Type:       "AWS::ApiGateway::Deployment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayDeployment parses ApiGatewayDeployment
func ParseApiGatewayDeployment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayDeployment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDeployment - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayDeployment validator
func (resource ApiGatewayDeployment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayDeploymentProperties validator
func (resource ApiGatewayDeploymentProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
