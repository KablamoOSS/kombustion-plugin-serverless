package properties

	import "fmt"

type InstanceFleetConfig_SpotProvisioningSpecification struct {
	
	
	
	BlockDurationMinutes interface{} `yaml:"BlockDurationMinutes,omitempty"`
	TimeoutAction interface{} `yaml:"TimeoutAction"`
	TimeoutDurationMinutes interface{} `yaml:"TimeoutDurationMinutes"`
}

func (resource InstanceFleetConfig_SpotProvisioningSpecification) Validate() []error {
	errs := []error{}
	
	
	
	if resource.TimeoutAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutAction'"))
	}
	if resource.TimeoutDurationMinutes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TimeoutDurationMinutes'"))
	}
	return errs
}
