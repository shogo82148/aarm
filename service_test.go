package aarm

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/shogo82148/hi"
	"github.com/shogo82148/hi/sets"
)

func testStructCompatibility(t *testing.T, a, b any, ignore []string) {
	typeA := reflect.TypeOf(a)
	setA := getFieldSet(typeA, ignore)
	typeB := reflect.TypeOf(b)
	setB := getFieldSet(typeB, ignore)
	for k := range setA {
		if !setB.Contains(k) {
			t.Errorf("%s has the field %s, but %s doesn't", typeA, k, typeB)
		}
	}
	for k := range setB {
		if !setA.Contains(k) {
			t.Errorf("%s has the field %s, but %s doesn't", typeB, k, typeA)
		}
	}
}

func getFieldSet(typ reflect.Type, ignore []string) sets.Set[string] {
	set := sets.New[string]()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.IsExported() {
			continue
		}
		if hi.Any(ignore, f.Name) {
			continue
		}
		set.Add(f.Name)
	}
	return set
}

func TestService(t *testing.T) {
	testStructCompatibility(t, Service{}, types.Service{}, []string{
		"AutoScalingConfigurationArn",
		"AutoScalingConfigurationSummary",
		"CreatedAt",
		"UpdatedAt",
		"DeletedAt",
		"Status",
		"ServiceArn",
		"ServiceId",
		"ServiceUrl",
	})
	testStructCompatibility(t, SourceConfiguration{}, types.SourceConfiguration{}, nil)
	testStructCompatibility(t, AuthenticationConfiguration{}, types.AuthenticationConfiguration{}, nil)
	testStructCompatibility(t, CodeRepository{}, types.CodeRepository{}, nil)
	testStructCompatibility(t, CodeConfiguration{}, types.CodeConfiguration{}, nil)
	testStructCompatibility(t, CodeConfigurationValues{}, types.CodeConfigurationValues{}, nil)
	testStructCompatibility(t, SourceCodeVersion{}, types.SourceCodeVersion{}, nil)
	testStructCompatibility(t, ImageRepository{}, types.ImageRepository{}, nil)
	testStructCompatibility(t, ImageConfiguration{}, types.ImageConfiguration{}, nil)
	testStructCompatibility(t, EncryptionConfiguration{}, types.EncryptionConfiguration{}, nil)
	testStructCompatibility(t, HealthCheckConfiguration{}, types.HealthCheckConfiguration{}, nil)
	testStructCompatibility(t, InstanceConfiguration{}, types.InstanceConfiguration{}, nil)
	testStructCompatibility(t, NetworkConfiguration{}, types.NetworkConfiguration{}, nil)
	testStructCompatibility(t, EgressConfiguration{}, types.EgressConfiguration{}, nil)
	testStructCompatibility(t, IngressConfiguration{}, types.IngressConfiguration{}, nil)
	testStructCompatibility(t, ServiceObservabilityConfiguration{}, types.ServiceObservabilityConfiguration{}, nil)
}
