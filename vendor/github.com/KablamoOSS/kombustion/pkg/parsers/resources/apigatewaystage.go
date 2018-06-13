package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayStage Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html
type ApiGatewayStage struct {
	Type       string                    `yaml:"Type"`
	Properties ApiGatewayStageProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// ApiGatewayStage Properties
type ApiGatewayStageProperties struct {
	CacheClusterEnabled  interface{} `yaml:"CacheClusterEnabled,omitempty"`
	CacheClusterSize     interface{} `yaml:"CacheClusterSize,omitempty"`
	ClientCertificateId  interface{} `yaml:"ClientCertificateId,omitempty"`
	DeploymentId         interface{} `yaml:"DeploymentId,omitempty"`
	Description          interface{} `yaml:"Description,omitempty"`
	DocumentationVersion interface{} `yaml:"DocumentationVersion,omitempty"`
	RestApiId            interface{} `yaml:"RestApiId"`
	StageName            interface{} `yaml:"StageName,omitempty"`
	Variables            interface{} `yaml:"Variables,omitempty"`
	MethodSettings       interface{} `yaml:"MethodSettings,omitempty"`
}

// NewApiGatewayStage constructor creates a new ApiGatewayStage
func NewApiGatewayStage(properties ApiGatewayStageProperties, deps ...interface{}) ApiGatewayStage {
	return ApiGatewayStage{
		Type:       "AWS::ApiGateway::Stage",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayStage parses ApiGatewayStage
func ParseApiGatewayStage(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayStage
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayStage - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayStage validator
func (resource ApiGatewayStage) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayStageProperties validator
func (resource ApiGatewayStageProperties) Validate() []error {
	errs := []error{}
	if resource.RestApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errs
}
