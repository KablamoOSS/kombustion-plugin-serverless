package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ConfigurationRecorderRecordingGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configurationrecorder-recordinggroup.html
type ConfigurationRecorderRecordingGroup struct {
	AllSupported               interface{} `yaml:"AllSupported,omitempty"`
	IncludeGlobalResourceTypes interface{} `yaml:"IncludeGlobalResourceTypes,omitempty"`
	ResourceTypes              interface{} `yaml:"ResourceTypes,omitempty"`
}

// ConfigurationRecorderRecordingGroup validation
func (resource ConfigurationRecorderRecordingGroup) Validate() []error {
	errs := []error{}

	return errs
}
