package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type PatchBaseline_Rule struct {
	ApproveAfterDays  interface{}                     `yaml:"ApproveAfterDays,omitempty"`
	ComplianceLevel   interface{}                     `yaml:"ComplianceLevel,omitempty"`
	EnableNonSecurity interface{}                     `yaml:"EnableNonSecurity,omitempty"`
	PatchFilterGroup  *PatchBaseline_PatchFilterGroup `yaml:"PatchFilterGroup,omitempty"`
}

func (resource PatchBaseline_Rule) Validate() []error {
	errs := []error{}

	return errs
}
