package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayAuthorizer Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
type ApiGatewayAuthorizer struct {
	Type       string                         `yaml:"Type"`
	Properties ApiGatewayAuthorizerProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// ApiGatewayAuthorizer Properties
type ApiGatewayAuthorizerProperties struct {
	AuthType                     interface{} `yaml:"AuthType,omitempty"`
	AuthorizerCredentials        interface{} `yaml:"AuthorizerCredentials,omitempty"`
	AuthorizerResultTtlInSeconds interface{} `yaml:"AuthorizerResultTtlInSeconds,omitempty"`
	AuthorizerUri                interface{} `yaml:"AuthorizerUri,omitempty"`
	IdentitySource               interface{} `yaml:"IdentitySource,omitempty"`
	IdentityValidationExpression interface{} `yaml:"IdentityValidationExpression,omitempty"`
	Name                         interface{} `yaml:"Name,omitempty"`
	RestApiId                    interface{} `yaml:"RestApiId"`
	Type                         interface{} `yaml:"Type,omitempty"`
	ProviderARNs                 interface{} `yaml:"ProviderARNs,omitempty"`
}

// NewApiGatewayAuthorizer constructor creates a new ApiGatewayAuthorizer
func NewApiGatewayAuthorizer(properties ApiGatewayAuthorizerProperties, deps ...interface{}) ApiGatewayAuthorizer {
	return ApiGatewayAuthorizer{
		Type:       "AWS::ApiGateway::Authorizer",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayAuthorizer parses ApiGatewayAuthorizer
func ParseApiGatewayAuthorizer(name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayAuthorizer
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayAuthorizer - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayAuthorizer validator
func (resource ApiGatewayAuthorizer) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayAuthorizerProperties validator
func (resource ApiGatewayAuthorizerProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
