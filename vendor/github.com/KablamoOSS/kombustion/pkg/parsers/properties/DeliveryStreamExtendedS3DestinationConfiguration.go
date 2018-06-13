package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// DeliveryStreamExtendedS3DestinationConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html
type DeliveryStreamExtendedS3DestinationConfiguration struct {
	BucketARN                interface{}                               `yaml:"BucketARN"`
	CompressionFormat        interface{}                               `yaml:"CompressionFormat"`
	Prefix                   interface{}                               `yaml:"Prefix"`
	RoleARN                  interface{}                               `yaml:"RoleARN"`
	S3BackupMode             interface{}                               `yaml:"S3BackupMode,omitempty"`
	S3BackupConfiguration    *DeliveryStreamS3DestinationConfiguration `yaml:"S3BackupConfiguration,omitempty"`
	ProcessingConfiguration  *DeliveryStreamProcessingConfiguration    `yaml:"ProcessingConfiguration,omitempty"`
	EncryptionConfiguration  *DeliveryStreamEncryptionConfiguration    `yaml:"EncryptionConfiguration,omitempty"`
	CloudWatchLoggingOptions *DeliveryStreamCloudWatchLoggingOptions   `yaml:"CloudWatchLoggingOptions,omitempty"`
	BufferingHints           *DeliveryStreamBufferingHints             `yaml:"BufferingHints"`
}

// DeliveryStreamExtendedS3DestinationConfiguration validation
func (resource DeliveryStreamExtendedS3DestinationConfiguration) Validate() []error {
	errs := []error{}

	if resource.BucketARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketARN'"))
	}
	if resource.CompressionFormat == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CompressionFormat'"))
	}
	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.BufferingHints == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BufferingHints'"))
	} else {
		errs = append(errs, resource.BufferingHints.Validate()...)
	}
	return errs
}
