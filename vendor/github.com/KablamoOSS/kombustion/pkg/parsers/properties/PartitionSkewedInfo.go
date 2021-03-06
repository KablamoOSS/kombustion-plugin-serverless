package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// PartitionSkewedInfo Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-skewedinfo.html
type PartitionSkewedInfo struct {
	SkewedColumnValueLocationMaps interface{} `yaml:"SkewedColumnValueLocationMaps,omitempty"`
	SkewedColumnNames             interface{} `yaml:"SkewedColumnNames,omitempty"`
	SkewedColumnValues            interface{} `yaml:"SkewedColumnValues,omitempty"`
}

// PartitionSkewedInfo validation
func (resource PartitionSkewedInfo) Validate() []error {
	errs := []error{}

	return errs
}
