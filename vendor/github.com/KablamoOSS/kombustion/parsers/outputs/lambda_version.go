package outputs
import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseLambdaVersion(name string, data string) (cf plugins.ValueMap, err error) {
	
	var resource, output plugins.ValueMap
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	
	cf = plugins.ValueMap{
		name: plugins.ValueMap{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-LambdaVersion-" + name,
				},
			},
		},
	}

	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Version"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-LambdaVersion-" + name + "-Version",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Version"] = output
	

	return
}
