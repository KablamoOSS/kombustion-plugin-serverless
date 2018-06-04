package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type KinesisFirehoseDeliveryStream struct {
	Type       string                                  `yaml:"Type"`
	Properties KinesisFirehoseDeliveryStreamProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

type KinesisFirehoseDeliveryStreamProperties struct {
	DeliveryStreamName                    interface{}                                                      `yaml:"DeliveryStreamName,omitempty"`
	DeliveryStreamType                    interface{}                                                      `yaml:"DeliveryStreamType,omitempty"`
	S3DestinationConfiguration            *properties.DeliveryStream_S3DestinationConfiguration            `yaml:"S3DestinationConfiguration,omitempty"`
	RedshiftDestinationConfiguration      *properties.DeliveryStream_RedshiftDestinationConfiguration      `yaml:"RedshiftDestinationConfiguration,omitempty"`
	KinesisStreamSourceConfiguration      *properties.DeliveryStream_KinesisStreamSourceConfiguration      `yaml:"KinesisStreamSourceConfiguration,omitempty"`
	ExtendedS3DestinationConfiguration    *properties.DeliveryStream_ExtendedS3DestinationConfiguration    `yaml:"ExtendedS3DestinationConfiguration,omitempty"`
	ElasticsearchDestinationConfiguration *properties.DeliveryStream_ElasticsearchDestinationConfiguration `yaml:"ElasticsearchDestinationConfiguration,omitempty"`
}

func NewKinesisFirehoseDeliveryStream(properties KinesisFirehoseDeliveryStreamProperties, deps ...interface{}) KinesisFirehoseDeliveryStream {
	return KinesisFirehoseDeliveryStream{
		Type:       "AWS::KinesisFirehose::DeliveryStream",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKinesisFirehoseDeliveryStream(name string, data string) (cf types.TemplateObject, err error) {
	var resource KinesisFirehoseDeliveryStream
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisFirehoseDeliveryStream - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource KinesisFirehoseDeliveryStream) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KinesisFirehoseDeliveryStreamProperties) Validate() []error {
	errs := []error{}
	return errs
}
