package outputs
import (
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseOpsWorksVolume(name string, data string) (cf plugins.ValueMap, err error) {
	
	cf = plugins.ValueMap{
		name: plugins.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-OpsWorksVolume-" + name,
				},
			},
		},
	}

	

	return
}
