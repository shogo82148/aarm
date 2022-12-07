package apprunneriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type DeploymentStarter interface {
	StartDeployment(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error)
}

type ServiceCreator interface {
	CreateService(ctx context.Context, params *apprunner.CreateServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.CreateServiceOutput, error)
}

type ServiceDescriber interface {
	DescribeService(ctx context.Context, params *apprunner.DescribeServiceInput, optFns ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error)
}

type ServicesLister interface {
	ListServices(ctx context.Context, params *apprunner.ListServicesInput, optFns ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error)
}
