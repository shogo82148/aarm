package aarm

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

type Service struct {
	ServiceName                 *string                            `json:",omitempty"`
	SourceConfiguration         *SourceConfiguration               `json:",omitempty"`
	AutoScalingConfigurationArn *string                            `json:",omitempty"`
	EncryptionConfiguration     *EncryptionConfiguration           `json:",omitempty"`
	HealthCheckConfiguration    *HealthCheckConfiguration          `json:",omitempty"`
	InstanceConfiguration       *InstanceConfiguration             `json:",omitempty"`
	NetworkConfiguration        *NetworkConfiguration              `json:",omitempty"`
	ObservabilityConfiguration  *ServiceObservabilityConfiguration `json:",omitempty"`
}

func importService(svc *types.Service) *Service {
	return &Service{
		ServiceName:                 svc.ServiceName,
		SourceConfiguration:         importSourceConfiguration(svc.SourceConfiguration),
		AutoScalingConfigurationArn: svc.AutoScalingConfigurationSummary.AutoScalingConfigurationArn,
		EncryptionConfiguration:     importEncryptionConfiguration(svc.EncryptionConfiguration),
		HealthCheckConfiguration:    importHealthCheckConfiguration(svc.HealthCheckConfiguration),
		InstanceConfiguration:       importInstanceConfiguration(svc.InstanceConfiguration),
		NetworkConfiguration:        importNetworkConfiguration(svc.NetworkConfiguration),
		ObservabilityConfiguration:  importServiceObservabilityConfiguration(svc.ObservabilityConfiguration),
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

func (v *SourceConfiguration) export() *types.SourceConfiguration {
	if v == nil {
		return nil
	}
	return &types.SourceConfiguration{
		AuthenticationConfiguration: v.AuthenticationConfiguration.export(),
		AutoDeploymentsEnabled:      v.AutoDeploymentsEnabled,
		CodeRepository:              v.CodeRepository.export(),
		ImageRepository:             v.ImageRepository.export(),
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

func (v *AuthenticationConfiguration) export() *types.AuthenticationConfiguration {
	if v == nil {
		return nil
	}
	return &types.AuthenticationConfiguration{
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

func (v *CodeRepository) export() *types.CodeRepository {
	if v == nil {
		return nil
	}
	return &types.CodeRepository{
		CodeConfiguration: v.CodeConfiguration.export(),
		RepositoryUrl:     v.RepositoryUrl,
		SourceCodeVersion: v.SourceCodeVersion.export(),
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

func (v *CodeConfiguration) export() *types.CodeConfiguration {
	if v == nil {
		return nil
	}
	return &types.CodeConfiguration{
		CodeConfigurationValues: v.CodeConfigurationValues.export(),
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

func (v *CodeConfigurationValues) export() *types.CodeConfigurationValues {
	if v == nil {
		return nil
	}
	return &types.CodeConfigurationValues{
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

func (v *SourceCodeVersion) export() *types.SourceCodeVersion {
	if v == nil {
		return nil
	}
	return &types.SourceCodeVersion{
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

func (v *ImageRepository) export() *types.ImageRepository {
	if v == nil {
		return nil
	}
	return &types.ImageRepository{
		ImageConfiguration:  v.ImageConfiguration.export(),
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

func (v *ImageConfiguration) export() *types.ImageConfiguration {
	if v == nil {
		return nil
	}
	return &types.ImageConfiguration{
		Port:                        v.Port,
		RuntimeEnvironmentVariables: v.RuntimeEnvironmentVariables,
		StartCommand:                v.StartCommand,
	}
}

type EncryptionConfiguration struct {
	KmsKey *string
}

func importEncryptionConfiguration(v *types.EncryptionConfiguration) *EncryptionConfiguration {
	if v == nil {
		return nil
	}
	return &EncryptionConfiguration{
		KmsKey: v.KmsKey,
	}
}

func (v *EncryptionConfiguration) export() *types.EncryptionConfiguration {
	if v == nil {
		return nil
	}
	return &types.EncryptionConfiguration{
		KmsKey: v.KmsKey,
	}
}

type HealthCheckConfiguration struct {
	HealthyThreshold   *int32
	Interval           *int32
	Path               *string
	Protocol           types.HealthCheckProtocol
	Timeout            *int32
	UnhealthyThreshold *int32
}

func importHealthCheckConfiguration(v *types.HealthCheckConfiguration) *HealthCheckConfiguration {
	if v == nil {
		return nil
	}
	return &HealthCheckConfiguration{
		HealthyThreshold:   v.HealthyThreshold,
		Interval:           v.Interval,
		Path:               v.Path,
		Protocol:           v.Protocol,
		Timeout:            v.Timeout,
		UnhealthyThreshold: v.UnhealthyThreshold,
	}
}

func (v *HealthCheckConfiguration) export() *types.HealthCheckConfiguration {
	if v == nil {
		return nil
	}
	return &types.HealthCheckConfiguration{
		HealthyThreshold:   v.HealthyThreshold,
		Interval:           v.Interval,
		Path:               v.Path,
		Protocol:           v.Protocol,
		Timeout:            v.Timeout,
		UnhealthyThreshold: v.UnhealthyThreshold,
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

func (v *InstanceConfiguration) export() *types.InstanceConfiguration {
	if v == nil {
		return nil
	}
	return &types.InstanceConfiguration{
		Cpu:             v.Cpu,
		InstanceRoleArn: v.InstanceRoleArn,
		Memory:          v.Memory,
	}
}

type NetworkConfiguration struct {
	// Network configuration settings for outbound message traffic.
	EgressConfiguration *EgressConfiguration `json:",omitempty"`

	// Network configuration settings for inbound message traffic.
	IngressConfiguration *IngressConfiguration `json:",omitempty"`
}

func importNetworkConfiguration(v *types.NetworkConfiguration) *NetworkConfiguration {
	if v == nil {
		return nil
	}
	return &NetworkConfiguration{
		EgressConfiguration:  importEgressConfiguration(v.EgressConfiguration),
		IngressConfiguration: importIngressConfiguration(v.IngressConfiguration),
	}
}

func (v *NetworkConfiguration) export() *types.NetworkConfiguration {
	if v == nil {
		return nil
	}
	return &types.NetworkConfiguration{
		EgressConfiguration:  v.EgressConfiguration.export(),
		IngressConfiguration: v.IngressConfiguration.export(),
	}
}

type EgressConfiguration struct {
	EgressType      types.EgressType
	VpcConnectorArn *string
}

func importEgressConfiguration(v *types.EgressConfiguration) *EgressConfiguration {
	if v == nil {
		return nil
	}
	return &EgressConfiguration{}
}

func (v *EgressConfiguration) export() *types.EgressConfiguration {
	return &types.EgressConfiguration{
		EgressType:      v.EgressType,
		VpcConnectorArn: v.VpcConnectorArn,
	}
}

type IngressConfiguration struct {
	IsPubliclyAccessible bool
}

func importIngressConfiguration(v *types.IngressConfiguration) *IngressConfiguration {
	if v == nil {
		return nil
	}
	return &IngressConfiguration{
		IsPubliclyAccessible: v.IsPubliclyAccessible,
	}
}

func (v *IngressConfiguration) export() *types.IngressConfiguration {
	if v == nil {
		return nil
	}
	return &types.IngressConfiguration{
		IsPubliclyAccessible: v.IsPubliclyAccessible,
	}
}

type ServiceObservabilityConfiguration struct {
	ObservabilityEnabled          bool
	ObservabilityConfigurationArn *string
}

func importServiceObservabilityConfiguration(v *types.ServiceObservabilityConfiguration) *ServiceObservabilityConfiguration {
	if v == nil {
		return nil
	}
	return &ServiceObservabilityConfiguration{
		ObservabilityEnabled:          v.ObservabilityEnabled,
		ObservabilityConfigurationArn: v.ObservabilityConfigurationArn,
	}
}

func (v *ServiceObservabilityConfiguration) export() *types.ServiceObservabilityConfiguration {
	if v == nil {
		return nil
	}
	return &types.ServiceObservabilityConfiguration{
		ObservabilityEnabled:          v.ObservabilityEnabled,
		ObservabilityConfigurationArn: v.ObservabilityConfigurationArn,
	}
}

// loadService load a service configuration from a file.
func (*App) loadService(filepath string) (*Service, error) {
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

func (*App) marshalService(svc any) ([]byte, error) {
	return json.MarshalIndent(svc, "", "  ")
}
