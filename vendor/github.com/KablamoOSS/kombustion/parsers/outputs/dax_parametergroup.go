package outputs
import (
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseDAXParameterGroup(name string, data string) (cf plugins.ValueMap, err error) {
	
	cf = plugins.ValueMap{
		name: plugins.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-DAXParameterGroup-" + name,
				},
			},
		},
	}

	

	return
}
