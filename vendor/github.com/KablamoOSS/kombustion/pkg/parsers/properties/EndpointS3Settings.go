package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// EndpointS3Settings Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html
type EndpointS3Settings struct {
	BucketFolder            interface{} `yaml:"BucketFolder,omitempty"`
	BucketName              interface{} `yaml:"BucketName,omitempty"`
	CompressionType         interface{} `yaml:"CompressionType,omitempty"`
	CsvDelimiter            interface{} `yaml:"CsvDelimiter,omitempty"`
	CsvRowDelimiter         interface{} `yaml:"CsvRowDelimiter,omitempty"`
	ExternalTableDefinition interface{} `yaml:"ExternalTableDefinition,omitempty"`
	ServiceAccessRoleArn    interface{} `yaml:"ServiceAccessRoleArn,omitempty"`
}

// EndpointS3Settings validation
func (resource EndpointS3Settings) Validate() []error {
	errs := []error{}

	return errs
}
