package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
	"log"
	"fmt"
)

type GuardDutyThreatIntelSet struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyThreatIntelSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyThreatIntelSetProperties struct {
	Activate interface{} `yaml:"Activate"`
	DetectorId interface{} `yaml:"DetectorId"`
	Format interface{} `yaml:"Format"`
	Location interface{} `yaml:"Location"`
	Name interface{} `yaml:"Name,omitempty"`
}

func NewGuardDutyThreatIntelSet(properties GuardDutyThreatIntelSetProperties, deps ...interface{}) GuardDutyThreatIntelSet {
	return GuardDutyThreatIntelSet{
		Type:       "AWS::GuardDuty::ThreatIntelSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyThreatIntelSet(name string, data string) (cf plugins.ValueMap, err error) {
	var resource GuardDutyThreatIntelSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyThreatIntelSet - ", err)
		}
		return
	}
	cf = plugins.ValueMap{name: resource}
	return
}

func (resource GuardDutyThreatIntelSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyThreatIntelSetProperties) Validate() []error {
	errs := []error{}
	if resource.Activate == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Activate'"))
	}
	if resource.DetectorId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Format == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Format'"))
	}
	if resource.Location == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Location'"))
	}
	return errs
}
