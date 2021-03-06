package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// JobDefinitionMountPoints Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html
type JobDefinitionMountPoints struct {
	ContainerPath interface{} `yaml:"ContainerPath,omitempty"`
	ReadOnly      interface{} `yaml:"ReadOnly,omitempty"`
	SourceVolume  interface{} `yaml:"SourceVolume,omitempty"`
}

// JobDefinitionMountPoints validation
func (resource JobDefinitionMountPoints) Validate() []error {
	errs := []error{}

	return errs
}
