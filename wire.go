//go:build wireinject
// +build wireinject

package aarm

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/google/wire"
	"github.com/shogo82148/aarm/internal/apprunneriface"
)

func NewApp(cfg aws.Config) *App {
	wire.Build(
		newApp, newAppRunner,
		wire.Struct(new(appRunner), "*"),
		wire.Bind(new(apprunneriface.DeploymentStarter), new(*apprunner.Client)),
		wire.Bind(new(apprunneriface.ServiceDescriber), new(*apprunner.Client)),
	)
	return &App{}
}
