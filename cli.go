package aarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

func CLI(ctx context.Context) (int, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return 1, err
	}

	svc := apprunner.NewFromConfig(cfg)
	app := &App{
		appRunner: &appRunner{
			DeploymentStarter: svc,
		},
	}
	app.Deploy(ctx)
	return 0, nil
}
