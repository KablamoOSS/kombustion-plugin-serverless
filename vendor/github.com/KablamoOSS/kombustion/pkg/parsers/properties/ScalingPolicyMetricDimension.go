package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPolicyMetricDimension Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-metricdimension.html
type ScalingPolicyMetricDimension struct {
	Name  interface{} `yaml:"Name"`
	Value interface{} `yaml:"Value"`
}

// ScalingPolicyMetricDimension validation
func (resource ScalingPolicyMetricDimension) Validate() []error {
	errs := []error{}

	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
