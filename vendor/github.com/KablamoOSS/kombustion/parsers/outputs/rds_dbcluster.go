package outputs
import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseRDSDBCluster(name string, data string) (cf plugins.ValueMap, err error) {
	
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
					"Fn::Sub": "${AWS::StackName}-RDSDBCluster-" + name,
				},
			},
		},
	}

	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Endpoint.Address"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-RDSDBCluster-" + name + "-EndpointAddress",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"EndpointAddress"] = output
	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Endpoint.Port"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-RDSDBCluster-" + name + "-EndpointPort",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"EndpointPort"] = output
	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "ReadEndpoint.Address"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-RDSDBCluster-" + name + "-ReadEndpointAddress",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"ReadEndpointAddress"] = output
	

	return
}
