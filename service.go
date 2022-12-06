package aarm

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

type Service struct {
	ServiceName           *string                `json:",omitempty"`
	InstanceConfiguration *InstanceConfiguration `json:",omitempty"`
	SourceConfiguration   *SourceConfiguration   `json:",omitempty"`
}

func importService(svc *types.Service) *Service {
	return &Service{
		ServiceName:           svc.ServiceName,
		InstanceConfiguration: importInstanceConfiguration(svc.InstanceConfiguration),
		SourceConfiguration:   importSourceConfiguration(svc.SourceConfiguration),
	}
}

type InstanceConfiguration struct {
	Cpu             *string `json:",omitempty"`
	InstanceRoleArn *string `json:",omitempty"`
	Memory          *string `json:",omitempty"`
}

func importInstanceConfiguration(v *types.InstanceConfiguration) *InstanceConfiguration {
	if v == nil {
		return nil
	}
	return &InstanceConfiguration{
		Cpu:             v.Cpu,
		InstanceRoleArn: v.InstanceRoleArn,
		Memory:          v.Memory,
	}
}

type SourceConfiguration struct {
	AuthenticationConfiguration *AuthenticationConfiguration `json:",omitempty"`
	AutoDeploymentsEnabled      *bool                        `json:",omitempty"`
	CodeRepository              *CodeRepository              `json:",omitempty"`
	ImageRepository             *ImageRepository             `json:",omitempty"`
}

func importSourceConfiguration(v *types.SourceConfiguration) *SourceConfiguration {
	if v == nil {
		return nil
	}
	return &SourceConfiguration{
		AuthenticationConfiguration: importAuthenticationConfiguration(v.AuthenticationConfiguration),
		AutoDeploymentsEnabled:      v.AutoDeploymentsEnabled,
		CodeRepository:              importCodeRepository(v.CodeRepository),
		ImageRepository:             importImageRepository(v.ImageRepository),
	}
}

type AuthenticationConfiguration struct {
	AccessRoleArn *string `json:",omitempty"`
	ConnectionArn *string `json:",omitempty"`
}

func importAuthenticationConfiguration(v *types.AuthenticationConfiguration) *AuthenticationConfiguration {
	if v == nil {
		return nil
	}
	return &AuthenticationConfiguration{
		AccessRoleArn: v.AccessRoleArn,
		ConnectionArn: v.ConnectionArn,
	}
}

type CodeRepository struct {
	CodeConfiguration *CodeConfiguration `json:",omitempty"`
	RepositoryUrl     *string            `json:",omitempty"`
	SourceCodeVersion *SourceCodeVersion `json:",omitempty"`
}

func importCodeRepository(v *types.CodeRepository) *CodeRepository {
	if v == nil {
		return nil
	}
	return &CodeRepository{
		CodeConfiguration: importCodeConfiguration(v.CodeConfiguration),
		RepositoryUrl:     v.RepositoryUrl,
		SourceCodeVersion: importSourceCodeVersion(v.SourceCodeVersion),
	}
}

type CodeConfiguration struct {
	CodeConfigurationValues *CodeConfigurationValues  `json:",omitempty"`
	ConfigurationSource     types.ConfigurationSource `json:",omitempty"`
}

func importCodeConfiguration(v *types.CodeConfiguration) *CodeConfiguration {
	if v == nil {
		return nil
	}
	return &CodeConfiguration{
		CodeConfigurationValues: importCodeConfigurationValues(v.CodeConfigurationValues),
		ConfigurationSource:     v.ConfigurationSource,
	}
}

type CodeConfigurationValues struct {
	BuildCommand                *string           `json:",omitempty"`
	Port                        *string           `json:",omitempty"`
	Runtime                     types.Runtime     `json:",omitempty"`
	RuntimeEnvironmentVariables map[string]string `json:",omitempty"`
	StartCommand                *string           `json:",omitempty"`
}

func importCodeConfigurationValues(v *types.CodeConfigurationValues) *CodeConfigurationValues {
	if v == nil {
		return nil
	}
	return &CodeConfigurationValues{
		BuildCommand:                v.BuildCommand,
		Port:                        v.Port,
		Runtime:                     v.Runtime,
		RuntimeEnvironmentVariables: v.RuntimeEnvironmentVariables,
		StartCommand:                v.StartCommand,
	}
}

type SourceCodeVersion struct {
	Type  types.SourceCodeVersionType `json:",omitempty"`
	Value *string                     `json:",omitempty"`
}

func importSourceCodeVersion(v *types.SourceCodeVersion) *SourceCodeVersion {
	if v == nil {
		return nil
	}
	return &SourceCodeVersion{
		Type:  v.Type,
		Value: v.Value,
	}
}

type ImageRepository struct {
	ImageConfiguration  *ImageConfiguration       `json:",omitempty"`
	ImageIdentifier     *string                   `json:",omitempty"`
	ImageRepositoryType types.ImageRepositoryType `json:",omitempty"`
}

func importImageRepository(v *types.ImageRepository) *ImageRepository {
	if v == nil {
		return nil
	}
	return &ImageRepository{
		ImageConfiguration:  importImageConfiguration(v.ImageConfiguration),
		ImageIdentifier:     v.ImageIdentifier,
		ImageRepositoryType: v.ImageRepositoryType,
	}
}

type ImageConfiguration struct {
	Port                        *string           `json:",omitempty"`
	RuntimeEnvironmentVariables map[string]string `json:",omitempty"`
	StartCommand                *string           `json:",omitempty"`
}

func importImageConfiguration(v *types.ImageConfiguration) *ImageConfiguration {
	if v == nil {
		return nil
	}
	return &ImageConfiguration{
		Port:                        v.Port,
		RuntimeEnvironmentVariables: v.RuntimeEnvironmentVariables,
		StartCommand:                v.StartCommand,
	}
}

// loadService load a service configuration from a file.
func loadService(filepath string) (*Service, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var srv Service
	if err := json.Unmarshal(data, &srv); err != nil {
		return nil, err
	}
	return &srv, nil
}
