package resources

import (
	"log"

	cfResources "github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	kombustionTypes "github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// LambdaPermissionConfig -
type LambdaPermissionConfig struct {
	Properties struct {
		Action           interface{} `yaml:"Action,omitempty"`
		EventSourceToken interface{} `yaml:"EventSourceToken,omitempty"`
		FunctionName     interface{} `yaml:"FunctionName"`
		Principal        interface{} `yaml:"Principal,omitempty"`
		SourceAccount    interface{} `yaml:"SourceAccount,omitempty"`
		SourceArn        interface{} `yaml:"SourceArn,omitempty"`
		SourceAPIGateway interface{} `yaml:"SourceApiGateway,omitempty"`
	} `yaml:"Properties"`
}

// ParseLambdaPermission -
func ParseLambdaPermission(name string, data string) (cf kombustionTypes.TemplateObject) {
	// Parse the config data
	var config LambdaPermissionConfig
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		log.Fatal(err)
		return
	}

	// validate the config
	config.Validate()

	action := config.Properties.Action
	if action == nil {
		action = "lambda:InvokeFunction"
	}

	principal := config.Properties.Principal
	if principal == nil {
		principal = "apigateway.amazonaws.com"
	}

	sourceArn := config.Properties.SourceArn
	if sourceArn == nil {
		if config.Properties.SourceAPIGateway != nil {
			sourceArn = map[string]interface{}{
				"Fn::Join": []interface{}{
					"", []interface{}{
						"arn:aws:execute-api:",
						map[string]string{"Ref": "AWS::Region"},
						":",
						map[string]string{"Ref": "AWS::AccountId"},
						":",
						config.Properties.SourceAPIGateway,
						"/*",
					},
				},
			}
		}
	}

	cf = kombustionTypes.TemplateObject{
		(name): cfResources.NewLambdaPermission(
			cfResources.LambdaPermissionProperties{
				Action:           action,
				Principal:        principal,
				SourceArn:        sourceArn,
				FunctionName:     config.Properties.FunctionName,
				SourceAccount:    config.Properties.SourceAccount,
			},
		),
	}
	return
}

// Validate - input Config validation
func (this LambdaPermissionConfig) Validate() {
	props := this.Properties
	if props.FunctionName == nil {
		log.Println("WARNING: LambdaPermissionConfig - Missing required field 'FunctionName'")
	}
	if props.SourceArn == nil && props.SourceAPIGateway == nil {
		log.Println("WARNING: LambdaPermissionConfig - Must provide one of: 'SourceArn', 'SourceApiGateway'")
	}
}
