package properties

	import "fmt"

type Instance_PrivateIpAddressSpecification struct {
	
	
	Primary interface{} `yaml:"Primary"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress"`
}

func (resource Instance_PrivateIpAddressSpecification) Validate() []error {
	errs := []error{}
	
	
	if resource.Primary == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Primary'"))
	}
	if resource.PrivateIpAddress == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PrivateIpAddress'"))
	}
	return errs
}
