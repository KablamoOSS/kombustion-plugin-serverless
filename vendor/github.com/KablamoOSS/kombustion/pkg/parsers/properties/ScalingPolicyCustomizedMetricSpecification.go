package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ScalingPolicyCustomizedMetricSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html
type ScalingPolicyCustomizedMetricSpecification struct {
	MetricName interface{} `yaml:"MetricName"`
	Namespace  interface{} `yaml:"Namespace"`
	Statistic  interface{} `yaml:"Statistic"`
	Unit       interface{} `yaml:"Unit,omitempty"`
	Dimensions interface{} `yaml:"Dimensions,omitempty"`
}

// ScalingPolicyCustomizedMetricSpecification validation
func (resource ScalingPolicyCustomizedMetricSpecification) Validate() []error {
	errs := []error{}

	if resource.MetricName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MetricName'"))
	}
	if resource.Namespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Namespace'"))
	}
	if resource.Statistic == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Statistic'"))
	}
	return errs
}
