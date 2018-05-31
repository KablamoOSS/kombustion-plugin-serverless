package outputs
import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/plugins"
)

func ParseElasticsearchDomain(name string, data string) (cf plugins.ValueMap, err error) {
	
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
					"Fn::Sub": "${AWS::StackName}-ElasticsearchDomain-" + name,
				},
			},
		},
	}

	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "Arn"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticsearchDomain-" + name + "-Arn",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"Arn"] = output
	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "DomainArn"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticsearchDomain-" + name + "-DomainArn",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"DomainArn"] = output
	
	output = plugins.ValueMap{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "DomainEndpoint"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-ElasticsearchDomain-" + name + "-DomainEndpoint",
			},
		},
	}
	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}
	cf[name+"DomainEndpoint"] = output
	

	return
}
