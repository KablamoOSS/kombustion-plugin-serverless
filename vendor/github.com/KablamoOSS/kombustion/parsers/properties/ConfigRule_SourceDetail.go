package properties

	import "fmt"

type ConfigRule_SourceDetail struct {
	
	
	
	EventSource interface{} `yaml:"EventSource"`
	MaximumExecutionFrequency interface{} `yaml:"MaximumExecutionFrequency,omitempty"`
	MessageType interface{} `yaml:"MessageType"`
}

func (resource ConfigRule_SourceDetail) Validate() []error {
	errs := []error{}
	
	
	
	if resource.EventSource == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EventSource'"))
	}
	if resource.MessageType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MessageType'"))
	}
	return errs
}
