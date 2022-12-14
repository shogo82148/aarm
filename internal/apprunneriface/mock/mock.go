package mock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type StartDeployment func(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error)

func (mock StartDeployment) StartDeployment(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error) {
	return mock(ctx, params, optFns...)
}

type CreateService func(ctx context.Context, params *apprunner.CreateServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.CreateServiceOutput, error)

func (mock CreateService) CreateService(ctx context.Context, params *apprunner.CreateServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.CreateServiceOutput, error) {
	return mock(ctx, params, optFns...)
}

type DescribeService func(ctx context.Context, params *apprunner.DescribeServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error)

func (mock DescribeService) DescribeService(ctx context.Context, params *apprunner.DescribeServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error) {
	return mock(ctx, params, optFns...)
}

type ListServices func(ctx context.Context, params *apprunner.ListServicesInput, optFns ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error)

func (mock ListServices) ListServices(ctx context.Context, params *apprunner.ListServicesInput, optFns ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error) {
	return mock(ctx, params, optFns...)
}
