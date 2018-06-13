package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineEncryptionKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html
type PipelineEncryptionKey struct {
	Id   interface{} `yaml:"Id"`
	Type interface{} `yaml:"Type"`
}

// PipelineEncryptionKey validation
func (resource PipelineEncryptionKey) Validate() []error {
	errs := []error{}

	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
