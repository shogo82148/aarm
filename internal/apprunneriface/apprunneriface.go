package apprunneriface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type DeploymentStarter interface {
	StartDeployment(ctx context.Context, params *apprunner.StartDeploymentInput, optFns ...func(*apprunner.Options)) (*apprunner.StartDeploymentOutput, error)
}
