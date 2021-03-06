package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CognitoUserPoolClient Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolclient.html
type CognitoUserPoolClient struct {
	Type       string                          `yaml:"Type"`
	Properties CognitoUserPoolClientProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// CognitoUserPoolClient Properties
type CognitoUserPoolClientProperties struct {
	ClientName           interface{} `yaml:"ClientName,omitempty"`
	GenerateSecret       interface{} `yaml:"GenerateSecret,omitempty"`
	RefreshTokenValidity interface{} `yaml:"RefreshTokenValidity,omitempty"`
	UserPoolId           interface{} `yaml:"UserPoolId"`
	ExplicitAuthFlows    interface{} `yaml:"ExplicitAuthFlows,omitempty"`
	ReadAttributes       interface{} `yaml:"ReadAttributes,omitempty"`
	WriteAttributes      interface{} `yaml:"WriteAttributes,omitempty"`
}

// NewCognitoUserPoolClient constructor creates a new CognitoUserPoolClient
func NewCognitoUserPoolClient(properties CognitoUserPoolClientProperties, deps ...interface{}) CognitoUserPoolClient {
	return CognitoUserPoolClient{
		Type:       "AWS::Cognito::UserPoolClient",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPoolClient parses CognitoUserPoolClient
func ParseCognitoUserPoolClient(name string, data string) (cf types.TemplateObject, err error) {
	var resource CognitoUserPoolClient
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPoolClient - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCognitoUserPoolClient validator
func (resource CognitoUserPoolClient) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoUserPoolClientProperties validator
func (resource CognitoUserPoolClientProperties) Validate() []error {
	errs := []error{}
	if resource.UserPoolId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UserPoolId'"))
	}
	return errs
}
