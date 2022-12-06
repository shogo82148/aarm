package aarm

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type DeployOption struct {
	GlobalOptions
	ServiceArn  string
	ServiceName string
	ConfigPath  string
}

func (opts *DeployOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
	set.StringVar(&opts.ConfigPath, "config-path", "service.json", "config path")
}

func (app *App) Deploy(ctx context.Context, opts *DeployOption) error {
	svc, err := loadService(opts.ConfigPath)
	if err != nil {
		return err
	}
	arn, err := app.getServiceArn(ctx, aws.ToString(svc.ServiceName))
	if err != nil {
		return err
	}
	log.Println(arn)
	_, err = app.appRunner.StartDeployment(ctx, &apprunner.StartDeploymentInput{
		ServiceArn: aws.String(arn),
	})
	if err != nil {
		return err
	}

	return nil
}

// getServiceArn finds the service that has name and return its arn.
func (app *App) getServiceArn(ctx context.Context, name string) (string, error) {
	paginator := apprunner.NewListServicesPaginator(app.appRunner, &apprunner.ListServicesInput{})
	for paginator.HasMorePages() {
		out, err := paginator.NextPage(ctx)
		if err != nil {
			return "", err
		}
		for _, s := range out.ServiceSummaryList {
			if aws.ToString(s.ServiceName) == name {
				return aws.ToString(s.ServiceArn), nil
			}
		}
	}
	return "", fmt.Errorf("aarm: service %q is not found", name)
}
