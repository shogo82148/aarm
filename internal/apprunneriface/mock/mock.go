package mock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type StartDeployment func(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error)

func (mock StartDeployment) StartDeployment(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error) {
	return mock(ctx, params, optFns...)
}
