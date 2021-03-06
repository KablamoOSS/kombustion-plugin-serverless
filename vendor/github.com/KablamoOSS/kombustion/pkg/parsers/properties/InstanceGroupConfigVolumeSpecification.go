package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// InstanceGroupConfigVolumeSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration-ebsblockdeviceconfig-volumespecification.html
type InstanceGroupConfigVolumeSpecification struct {
	Iops       interface{} `yaml:"Iops,omitempty"`
	SizeInGB   interface{} `yaml:"SizeInGB"`
	VolumeType interface{} `yaml:"VolumeType"`
}

// InstanceGroupConfigVolumeSpecification validation
func (resource InstanceGroupConfigVolumeSpecification) Validate() []error {
	errs := []error{}

	if resource.SizeInGB == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SizeInGB'"))
	}
	if resource.VolumeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeType'"))
	}
	return errs
}
