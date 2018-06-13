package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TableGlobalSecondaryIndex Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html
type TableGlobalSecondaryIndex struct {
	IndexName             interface{}                 `yaml:"IndexName"`
	ProvisionedThroughput *TableProvisionedThroughput `yaml:"ProvisionedThroughput"`
	Projection            *TableProjection            `yaml:"Projection"`
	KeySchema             interface{}                 `yaml:"KeySchema"`
}

// TableGlobalSecondaryIndex validation
func (resource TableGlobalSecondaryIndex) Validate() []error {
	errs := []error{}

	if resource.IndexName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IndexName'"))
	}
	if resource.ProvisionedThroughput == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ProvisionedThroughput'"))
	} else {
		errs = append(errs, resource.ProvisionedThroughput.Validate()...)
	}
	if resource.Projection == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Projection'"))
	} else {
		errs = append(errs, resource.Projection.Validate()...)
	}
	if resource.KeySchema == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeySchema'"))
	}
	return errs
}
