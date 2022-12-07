package aarm

import (
	"context"
	"flag"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type InitOption struct {
	GlobalOptions
	ServiceArn  string
	ServiceName string
}

func (opts *InitOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
	set.StringVar(&opts.ServiceArn, "service-arn", "", "service arn")
	set.StringVar(&opts.ServiceName, "service-name", "", "service name")
}

func (app *App) Init(ctx context.Context, opts *InitOption) error {
	out, err := app.appRunner.DescribeService(ctx, &apprunner.DescribeServiceInput{
		ServiceArn: aws.String(opts.ServiceArn),
	})
	if err != nil {
		return err
	}

	service := importService(out.Service)
	data, err := app.marshalService(service)
	if err != nil {
		return err
	}
	return os.WriteFile(app.cfg.Service, data, 0o644)
}
