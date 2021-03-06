package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StreamingDistributionTrustedSigners Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-trustedsigners.html
type StreamingDistributionTrustedSigners struct {
	Enabled           interface{} `yaml:"Enabled"`
	AwsAccountNumbers interface{} `yaml:"AwsAccountNumbers,omitempty"`
}

// StreamingDistributionTrustedSigners validation
func (resource StreamingDistributionTrustedSigners) Validate() []error {
	errs := []error{}

	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	return errs
}
