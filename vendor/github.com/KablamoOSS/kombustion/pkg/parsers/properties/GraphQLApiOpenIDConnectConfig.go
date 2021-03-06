package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// GraphQLApiOpenIDConnectConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html
type GraphQLApiOpenIDConnectConfig struct {
	AuthTTL  interface{} `yaml:"AuthTTL,omitempty"`
	ClientId interface{} `yaml:"ClientId,omitempty"`
	IatTTL   interface{} `yaml:"IatTTL,omitempty"`
	Issuer   interface{} `yaml:"Issuer,omitempty"`
}

// GraphQLApiOpenIDConnectConfig validation
func (resource GraphQLApiOpenIDConnectConfig) Validate() []error {
	errs := []error{}

	return errs
}
