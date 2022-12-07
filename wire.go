//go:build wireinject
// +build wireinject

package aarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/google/wire"
	"github.com/shogo82148/aarm/internal/apprunneriface"
)

func NewApp(ctx context.Context, opts *GlobalOptions) (*App, error) {
	wire.Build(
		newApp, newAppRunner, newAWSConfig,
		wire.Struct(new(appRunner), "*"),
		wire.Bind(new(apprunneriface.DeploymentStarter), new(*apprunner.Client)),
		wire.Bind(new(apprunneriface.ServiceCreator), new(*apprunner.Client)),
		wire.Bind(new(apprunneriface.ServiceDescriber), new(*apprunner.Client)),
		wire.Bind(new(apprunneriface.ServicesLister), new(*apprunner.Client)),
	)
	return &App{}, nil
}
