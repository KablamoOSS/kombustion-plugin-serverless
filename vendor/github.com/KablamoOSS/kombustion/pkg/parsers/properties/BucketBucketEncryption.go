package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketBucketEncryption Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-bucketencryption.html
type BucketBucketEncryption struct {
	ServerSideEncryptionConfiguration interface{} `yaml:"ServerSideEncryptionConfiguration"`
}

// BucketBucketEncryption validation
func (resource BucketBucketEncryption) Validate() []error {
	errs := []error{}

	if resource.ServerSideEncryptionConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServerSideEncryptionConfiguration'"))
	}
	return errs
}
