package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApiKeyStageKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-apikey-stagekey.html
type ApiKeyStageKey struct {
	RestApiId interface{} `yaml:"RestApiId,omitempty"`
	StageName interface{} `yaml:"StageName,omitempty"`
}

// ApiKeyStageKey validation
func (resource ApiKeyStageKey) Validate() []error {
	errs := []error{}

	return errs
}
