package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClusterSimpleScalingPolicyConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html
type ClusterSimpleScalingPolicyConfiguration struct {
	AdjustmentType    interface{} `yaml:"AdjustmentType,omitempty"`
	CoolDown          interface{} `yaml:"CoolDown,omitempty"`
	ScalingAdjustment interface{} `yaml:"ScalingAdjustment"`
}

// ClusterSimpleScalingPolicyConfiguration validation
func (resource ClusterSimpleScalingPolicyConfiguration) Validate() []error {
	errs := []error{}

	if resource.ScalingAdjustment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalingAdjustment'"))
	}
	return errs
}
