package outputs
import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseElasticBeanstalkEnvironment(name string, data string) (cf plugins.ValueMap, err error) {
	
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
					"Fn::Sub": "${AWS::StackName}-ElasticBeanstalkEnvironment-" + name,
				},
			},
		},
	}

	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "EndpointURL"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticBeanstalkEnvironment-" + name + "-EndpointURL",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"EndpointURL"] = output
	

	return
}
