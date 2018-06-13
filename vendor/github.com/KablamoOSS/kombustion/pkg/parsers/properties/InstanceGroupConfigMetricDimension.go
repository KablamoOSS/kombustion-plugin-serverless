package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceGroupConfigMetricDimension Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html
type InstanceGroupConfigMetricDimension struct {
	Key   interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

// InstanceGroupConfigMetricDimension validation
func (resource InstanceGroupConfigMetricDimension) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
