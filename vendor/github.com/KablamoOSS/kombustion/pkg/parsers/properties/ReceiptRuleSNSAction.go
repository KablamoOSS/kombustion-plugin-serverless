package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ReceiptRuleSNSAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-snsaction.html
type ReceiptRuleSNSAction struct {
	Encoding interface{} `yaml:"Encoding,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

// ReceiptRuleSNSAction validation
func (resource ReceiptRuleSNSAction) Validate() []error {
	errs := []error{}

	return errs
}
