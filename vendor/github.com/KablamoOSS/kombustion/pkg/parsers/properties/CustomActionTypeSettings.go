package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// CustomActionTypeSettings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html
type CustomActionTypeSettings struct {
	EntityUrlTemplate          interface{} `yaml:"EntityUrlTemplate,omitempty"`
	ExecutionUrlTemplate       interface{} `yaml:"ExecutionUrlTemplate,omitempty"`
	RevisionUrlTemplate        interface{} `yaml:"RevisionUrlTemplate,omitempty"`
	ThirdPartyConfigurationUrl interface{} `yaml:"ThirdPartyConfigurationUrl,omitempty"`
}

// CustomActionTypeSettings validation
func (resource CustomActionTypeSettings) Validate() []error {
	errs := []error{}

	return errs
}
