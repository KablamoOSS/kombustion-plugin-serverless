package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// FunctionDeadLetterConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-deadletterconfig.html
type FunctionDeadLetterConfig struct {
	TargetArn interface{} `yaml:"TargetArn,omitempty"`
}

// FunctionDeadLetterConfig validation
func (resource FunctionDeadLetterConfig) Validate() []error {
	errs := []error{}

	return errs
}
