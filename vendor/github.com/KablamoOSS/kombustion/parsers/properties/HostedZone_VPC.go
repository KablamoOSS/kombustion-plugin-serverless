package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

type HostedZone_VPC struct {
	VPCId     interface{} `yaml:"VPCId"`
	VPCRegion interface{} `yaml:"VPCRegion"`
}

func (resource HostedZone_VPC) Validate() []error {
	errs := []error{}

	if resource.VPCId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VPCId'"))
	}
	if resource.VPCRegion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VPCRegion'"))
	}
	return errs
}
