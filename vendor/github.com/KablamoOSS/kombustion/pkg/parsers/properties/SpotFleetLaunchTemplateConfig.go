package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SpotFleetLaunchTemplateConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-launchtemplateconfig.html
type SpotFleetLaunchTemplateConfig struct {
	Overrides                   interface{}                                `yaml:"Overrides,omitempty"`
	LaunchTemplateSpecification *SpotFleetFleetLaunchTemplateSpecification `yaml:"LaunchTemplateSpecification,omitempty"`
}

// SpotFleetLaunchTemplateConfig validation
func (resource SpotFleetLaunchTemplateConfig) Validate() []error {
	errs := []error{}

	return errs
}
