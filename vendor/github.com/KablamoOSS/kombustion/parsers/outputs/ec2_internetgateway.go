package outputs
import (
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseEC2InternetGateway(name string, data string) (cf plugins.ValueMap, err error) {
	
	cf = plugins.ValueMap{
		name: plugins.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EC2InternetGateway-" + name,
				},
			},
		},
	}

	

	return
}
