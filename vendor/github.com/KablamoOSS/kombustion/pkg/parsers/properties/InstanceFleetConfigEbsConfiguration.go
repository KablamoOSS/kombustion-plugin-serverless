package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceFleetConfigEbsConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-ebsconfiguration.html
type InstanceFleetConfigEbsConfiguration struct {
	EbsOptimized          interface{} `yaml:"EbsOptimized,omitempty"`
	EbsBlockDeviceConfigs interface{} `yaml:"EbsBlockDeviceConfigs,omitempty"`
}

// InstanceFleetConfigEbsConfiguration validation
func (resource InstanceFleetConfigEbsConfiguration) Validate() []error {
	errs := []error{}

	return errs
}
