package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type Bucket_BucketEncryption struct {
	ServerSideEncryptionConfiguration interface{} `yaml:"ServerSideEncryptionConfiguration"`
}

func (resource Bucket_BucketEncryption) Validate() []error {
	errs := []error{}

	if resource.ServerSideEncryptionConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServerSideEncryptionConfiguration'"))
	}
	return errs
}
