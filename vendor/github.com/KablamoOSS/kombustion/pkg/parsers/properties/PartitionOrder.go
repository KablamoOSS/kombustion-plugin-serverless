package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PartitionOrder Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-order.html
type PartitionOrder struct {
	Column    interface{} `yaml:"Column"`
	SortOrder interface{} `yaml:"SortOrder,omitempty"`
}

// PartitionOrder validation
func (resource PartitionOrder) Validate() []error {
	errs := []error{}

	if resource.Column == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Column'"))
	}
	return errs
}
