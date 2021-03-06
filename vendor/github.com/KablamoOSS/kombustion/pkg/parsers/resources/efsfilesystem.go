package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EFSFileSystem Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
type EFSFileSystem struct {
	Type       string                  `yaml:"Type"`
	Properties EFSFileSystemProperties `yaml:"Properties"`
	Condition  interface{}             `yaml:"Condition,omitempty"`
	Metadata   interface{}             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}             `yaml:"DependsOn,omitempty"`
}

// EFSFileSystem Properties
type EFSFileSystemProperties struct {
	Encrypted       interface{} `yaml:"Encrypted,omitempty"`
	KmsKeyId        interface{} `yaml:"KmsKeyId,omitempty"`
	PerformanceMode interface{} `yaml:"PerformanceMode,omitempty"`
	FileSystemTags  interface{} `yaml:"FileSystemTags,omitempty"`
}

// NewEFSFileSystem constructor creates a new EFSFileSystem
func NewEFSFileSystem(properties EFSFileSystemProperties, deps ...interface{}) EFSFileSystem {
	return EFSFileSystem{
		Type:       "AWS::EFS::FileSystem",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEFSFileSystem parses EFSFileSystem
func ParseEFSFileSystem(name string, data string) (cf types.TemplateObject, err error) {
	var resource EFSFileSystem
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EFSFileSystem - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseEFSFileSystem validator
func (resource EFSFileSystem) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEFSFileSystemProperties validator
func (resource EFSFileSystemProperties) Validate() []error {
	errs := []error{}
	return errs
}
