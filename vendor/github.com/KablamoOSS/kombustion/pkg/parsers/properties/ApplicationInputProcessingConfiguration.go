package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ApplicationInputProcessingConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputprocessingconfiguration.html
type ApplicationInputProcessingConfiguration struct {
	InputLambdaProcessor *ApplicationInputLambdaProcessor `yaml:"InputLambdaProcessor,omitempty"`
}

// ApplicationInputProcessingConfiguration validation
func (resource ApplicationInputProcessingConfiguration) Validate() []error {
	errs := []error{}

	return errs
}