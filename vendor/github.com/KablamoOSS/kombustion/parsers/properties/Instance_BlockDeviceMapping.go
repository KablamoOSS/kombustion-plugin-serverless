package properties


type Instance_BlockDeviceMapping struct {
	
	
	
	
	DeviceName interface{} `yaml:"DeviceName,omitempty"`
	NoDevice interface{} `yaml:"NoDevice,omitempty"`
	VirtualName interface{} `yaml:"VirtualName,omitempty"`
	Ebs *Instance_EbsBlockDevice `yaml:"Ebs,omitempty"`
}

func (resource Instance_BlockDeviceMapping) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
