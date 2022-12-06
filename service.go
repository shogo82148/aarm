package aarm

import "github.com/aws/aws-sdk-go-v2/service/apprunner/types"

type Service struct {
	ServiceName           *string                `json:",omitempty"`
	InstanceConfiguration *InstanceConfiguration `json:",omitempty"`
	SourceConfiguration   *SourceConfiguration   `json:",omitempty"`
}

type InstanceConfiguration struct {
	Cpu             *string `json:",omitempty"`
	InstanceRoleArn *string `json:",omitempty"`
	Memory          *string `json:",omitempty"`
}

type SourceConfiguration struct {
	AuthenticationConfiguration *AuthenticationConfiguration `json:",omitempty"`
	AutoDeploymentsEnabled      *bool                        `json:",omitempty"`
	CodeRepository              *CodeRepository              `json:",omitempty"`
	ImageRepository             *ImageRepository             `json:",omitempty"`
}

type AuthenticationConfiguration struct {
	AccessRoleArn *string `json:",omitempty"`
	ConnectionArn *string `json:",omitempty"`
}

type CodeRepository struct {
	CodeConfiguration *CodeConfiguration `json:",omitempty"`
	RepositoryUrl     *string            `json:",omitempty"`
	SourceCodeVersion *SourceCodeVersion
}

type CodeConfiguration struct {
	CodeConfigurationValues *CodeConfigurationValues `json:",omitempty"`
	ConfigurationSource     *string                  `json:",omitempty"`
}

type CodeConfigurationValues struct {
	BuildCommand                *string         `json:",omitempty"`
	Port                        *string         `json:",omitempty"`
	Runtime                     *string         `json:",omitempty"`
	RuntimeEnvironmentVariables []*KeyValuePair `json:",omitempty"`
	StartCommand                *string         `json:",omitempty"`
}

type SourceCodeVersion struct {
	Type  *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

type ImageRepository struct {
	ImageConfiguration  *ImageConfiguration `json:",omitempty"`
	ImageIdentifier     *string             `json:",omitempty"`
	ImageRepositoryType *string             `json:",omitempty"`
}

type ImageConfiguration struct {
	Port                        *string         `json:",omitempty"`
	RuntimeEnvironmentVariables []*KeyValuePair `json:",omitempty"`
	StartCommand                *string         `json:",omitempty"`
}

type KeyValuePair struct {
	Name  *string `json:",omitempty"`
	Value *string `json:",omitempty"`
}

func importService(svc *types.Service) *Service {
	svcSrc := svc.SourceConfiguration
	src := &SourceConfiguration{
		AutoDeploymentsEnabled: svcSrc.AutoDeploymentsEnabled,
	}
	if auth := svcSrc.AuthenticationConfiguration; auth != nil {
		src.AuthenticationConfiguration = &AuthenticationConfiguration{
			AccessRoleArn: auth.AccessRoleArn,
			ConnectionArn: auth.ConnectionArn,
		}
	}
	return &Service{
		ServiceName: svc.ServiceName,
		InstanceConfiguration: &InstanceConfiguration{
			Cpu:             svc.InstanceConfiguration.Cpu,
			InstanceRoleArn: svc.InstanceConfiguration.InstanceRoleArn,
			Memory:          svc.InstanceConfiguration.Memory,
		},
		SourceConfiguration: src,
	}
}
