package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

type CertificateManagerCertificate struct {
	Type       string                                  `yaml:"Type"`
	Properties CertificateManagerCertificateProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

type CertificateManagerCertificateProperties struct {
	DomainName              interface{} `yaml:"DomainName"`
	DomainValidationOptions interface{} `yaml:"DomainValidationOptions,omitempty"`
	SubjectAlternativeNames interface{} `yaml:"SubjectAlternativeNames,omitempty"`
	Tags                    interface{} `yaml:"Tags,omitempty"`
}

func NewCertificateManagerCertificate(properties CertificateManagerCertificateProperties, deps ...interface{}) CertificateManagerCertificate {
	return CertificateManagerCertificate{
		Type:       "AWS::CertificateManager::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCertificateManagerCertificate(name string, data string) (cf types.TemplateObject, err error) {
	var resource CertificateManagerCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CertificateManagerCertificate - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

func (resource CertificateManagerCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CertificateManagerCertificateProperties) Validate() []error {
	errs := []error{}
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errs
}
