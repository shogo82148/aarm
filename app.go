package aarm

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/shogo82148/aarm/internal/apprunneriface"
)

type appRunner struct {
	apprunneriface.DeploymentStarter
	apprunneriface.ServiceDescriber
	apprunneriface.ServicesLister
}

type App struct {
	appRunner *appRunner
}

func newApp(runner *appRunner) *App {
	return &App{
		appRunner: runner,
	}
}

func newAppRunner(cfg aws.Config) *apprunner.Client {
	return apprunner.NewFromConfig(cfg)
}
