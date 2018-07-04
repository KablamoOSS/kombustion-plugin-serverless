package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CognitoUserPool Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpool.html
type CognitoUserPool struct {
	Type       string                    `yaml:"Type"`
	Properties CognitoUserPoolProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// CognitoUserPool Properties
type CognitoUserPoolProperties struct {
	EmailVerificationMessage interface{}                               `yaml:"EmailVerificationMessage,omitempty"`
	EmailVerificationSubject interface{}                               `yaml:"EmailVerificationSubject,omitempty"`
	MfaConfiguration         interface{}                               `yaml:"MfaConfiguration,omitempty"`
	SmsAuthenticationMessage interface{}                               `yaml:"SmsAuthenticationMessage,omitempty"`
	SmsVerificationMessage   interface{}                               `yaml:"SmsVerificationMessage,omitempty"`
	UserPoolName             interface{}                               `yaml:"UserPoolName,omitempty"`
	UserPoolTags             interface{}                               `yaml:"UserPoolTags,omitempty"`
	SmsConfiguration         *properties.UserPoolSmsConfiguration      `yaml:"SmsConfiguration,omitempty"`
	Policies                 *properties.UserPoolPolicies              `yaml:"Policies,omitempty"`
	AliasAttributes          interface{}                               `yaml:"AliasAttributes,omitempty"`
	UsernameAttributes       interface{}                               `yaml:"UsernameAttributes,omitempty"`
	Schema                   interface{}                               `yaml:"Schema,omitempty"`
	AutoVerifiedAttributes   interface{}                               `yaml:"AutoVerifiedAttributes,omitempty"`
	LambdaConfig             *properties.UserPoolLambdaConfig          `yaml:"LambdaConfig,omitempty"`
	EmailConfiguration       *properties.UserPoolEmailConfiguration    `yaml:"EmailConfiguration,omitempty"`
	DeviceConfiguration      *properties.UserPoolDeviceConfiguration   `yaml:"DeviceConfiguration,omitempty"`
	AdminCreateUserConfig    *properties.UserPoolAdminCreateUserConfig `yaml:"AdminCreateUserConfig,omitempty"`
}

// NewCognitoUserPool constructor creates a new CognitoUserPool
func NewCognitoUserPool(properties CognitoUserPoolProperties, deps ...interface{}) CognitoUserPool {
	return CognitoUserPool{
		Type:       "AWS::Cognito::UserPool",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPool parses CognitoUserPool
func ParseCognitoUserPool(name string, data string) (cf types.TemplateObject, err error) {
	var resource CognitoUserPool
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CognitoUserPool - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCognitoUserPool validator
func (resource CognitoUserPool) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoUserPoolProperties validator
func (resource CognitoUserPoolProperties) Validate() []error {
	errs := []error{}
	return errs
}
