package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PartitionSerdeInfo Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-serdeinfo.html
type PartitionSerdeInfo struct {
	Name                 interface{} `yaml:"Name,omitempty"`
	Parameters           interface{} `yaml:"Parameters,omitempty"`
	SerializationLibrary interface{} `yaml:"SerializationLibrary,omitempty"`
}

// PartitionSerdeInfo validation
func (resource PartitionSerdeInfo) Validate() []error {
	errs := []error{}

	return errs
}
