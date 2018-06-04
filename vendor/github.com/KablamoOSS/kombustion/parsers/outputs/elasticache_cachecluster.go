package outputs

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

func ParseElastiCacheCacheCluster(name string, data string) (cf types.TemplateObject, err error) {

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
					"Fn::Sub": "${AWS::StackName}-ElastiCacheCacheCluster-" + name,
				},
			},
		},
	}

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "ConfigurationEndpoint.Address"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElastiCacheCacheCluster-" + name + "-ConfigurationEndpointAddress",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"ConfigurationEndpointAddress"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "ConfigurationEndpoint.Port"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElastiCacheCacheCluster-" + name + "-ConfigurationEndpointPort",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"ConfigurationEndpointPort"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "RedisEndpoint.Address"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElastiCacheCacheCluster-" + name + "-RedisEndpointAddress",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"RedisEndpointAddress"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "RedisEndpoint.Port"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElastiCacheCacheCluster-" + name + "-RedisEndpointPort",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"RedisEndpointPort"] = output

	return
}
