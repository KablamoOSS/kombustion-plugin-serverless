package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayRequestValidator Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-requestvalidator.html
type ApiGatewayRequestValidator struct {
	Type       string                               `yaml:"Type"`
	Properties ApiGatewayRequestValidatorProperties `yaml:"Properties"`
	Condition  interface{}                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                          `yaml:"DependsOn,omitempty"`
}

// ApiGatewayRequestValidator Properties
type ApiGatewayRequestValidatorProperties struct {
	Name                      interface{} `yaml:"Name,omitempty"`
	RestApiId                 interface{} `yaml:"RestApiId"`
	ValidateRequestBody       interface{} `yaml:"ValidateRequestBody,omitempty"`
	ValidateRequestParameters interface{} `yaml:"ValidateRequestParameters,omitempty"`
}

// NewApiGatewayRequestValidator constructor creates a new ApiGatewayRequestValidator
func NewApiGatewayRequestValidator(properties ApiGatewayRequestValidatorProperties, deps ...interface{}) ApiGatewayRequestValidator {
	return ApiGatewayRequestValidator{
		Type:       "AWS::ApiGateway::RequestValidator",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayRequestValidator parses ApiGatewayRequestValidator
func ParseApiGatewayRequestValidator(name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayRequestValidator
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayRequestValidator - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayRequestValidator validator
func (resource ApiGatewayRequestValidator) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayRequestValidatorProperties validator
func (resource ApiGatewayRequestValidatorProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
