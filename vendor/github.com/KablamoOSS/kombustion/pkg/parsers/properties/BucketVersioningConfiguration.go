package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketVersioningConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type BucketVersioningConfiguration struct {
	Status interface{} `yaml:"Status"`
}

// BucketVersioningConfiguration validation
func (resource BucketVersioningConfiguration) Validate() []error {
	errs := []error{}

	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	return errs
}