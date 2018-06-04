package outputs

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

func ParseRDSDBCluster(name string, data string) (cf types.TemplateObject, err error) {

	var resource, output types.TemplateObject
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}

	cf = types.TemplateObject{
		name: types.TemplateObject{
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

	output = types.TemplateObject{
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

	output = types.TemplateObject{
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

	output = types.TemplateObject{
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
