package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ReceiptRuleBounceAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-bounceaction.html
type ReceiptRuleBounceAction struct {
	Message       interface{} `yaml:"Message"`
	Sender        interface{} `yaml:"Sender"`
	SmtpReplyCode interface{} `yaml:"SmtpReplyCode"`
	StatusCode    interface{} `yaml:"StatusCode,omitempty"`
	TopicArn      interface{} `yaml:"TopicArn,omitempty"`
}

// ReceiptRuleBounceAction validation
func (resource ReceiptRuleBounceAction) Validate() []error {
	errs := []error{}

	if resource.Message == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Message'"))
	}
	if resource.Sender == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Sender'"))
	}
	if resource.SmtpReplyCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SmtpReplyCode'"))
	}
	return errs
}
