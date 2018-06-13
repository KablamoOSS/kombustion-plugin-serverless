package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceBlockDeviceMapping Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-blockdevicemapping.html
type InstanceBlockDeviceMapping struct {
	DeviceName  interface{}             `yaml:"DeviceName,omitempty"`
	NoDevice    interface{}             `yaml:"NoDevice,omitempty"`
	VirtualName interface{}             `yaml:"VirtualName,omitempty"`
	Ebs         *InstanceEbsBlockDevice `yaml:"Ebs,omitempty"`
}

// InstanceBlockDeviceMapping validation
func (resource InstanceBlockDeviceMapping) Validate() []error {
	errs := []error{}

	return errs
}