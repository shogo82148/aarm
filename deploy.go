package aarm

import (
	"context"
	"flag"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type DeployOption struct {
	GlobalOptions
}

func (opts *DeployOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
}

func (app *App) Deploy(ctx context.Context, opts *DeployOption) error {
	svc, err := app.loadService()
	if err != nil {
		return err
	}

	// find the service from the list
	arn, err := app.getServiceArn(ctx, aws.ToString(svc.ServiceName))
	if err != nil {
		if _, ok := as[*serviceNotFoundError](err); ok {
			// service not found. create a new one.
			app.Log("[INFO] create a new service %q", aws.ToString(svc.ServiceName))
			if err := app.createService(ctx, svc); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	app.Log("[DEBUG] describe %q", arn)
	out, err := app.appRunner.DescribeService(ctx, &apprunner.DescribeServiceInput{
		ServiceArn: aws.String(arn),
	})
	if err != nil {
		return nil
	}
	remote := importService(out.Service)

	if reflect.DeepEqual(svc, remote) {
		// no need to update. start new deployment.
		app.Log("[INFO] start new deployment on %q", arn)
		_, err = app.appRunner.StartDeployment(ctx, &apprunner.StartDeploymentInput{
			ServiceArn: aws.String(arn),
		})
		if err != nil {
			return err
		}
		return nil
	}

	// need to update
	app.Log("[INFO] update service on %q", arn)
	err = app.updateService(ctx, arn, svc)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) createService(ctx context.Context, svc *Service) error {
	app.Log("[DEBUG] create service %q", aws.ToString(svc.ServiceName))
	_, err := app.appRunner.CreateService(ctx, &apprunner.CreateServiceInput{
		ServiceName:                 svc.ServiceName,
		SourceConfiguration:         svc.SourceConfiguration.export(),
		AutoScalingConfigurationArn: svc.AutoScalingConfigurationArn,
		EncryptionConfiguration:     svc.EncryptionConfiguration.export(),
		HealthCheckConfiguration:    svc.HealthCheckConfiguration.export(),
		InstanceConfiguration:       svc.InstanceConfiguration.export(),
		NetworkConfiguration:        svc.NetworkConfiguration.export(),
		ObservabilityConfiguration:  svc.ObservabilityConfiguration.export(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (app *App) updateService(ctx context.Context, arn string, svc *Service) error {
	app.Log("[DEBUG] update service on %q", arn)
	_, err := app.appRunner.UpdateService(ctx, &apprunner.UpdateServiceInput{
		ServiceArn:                  aws.String(arn),
		SourceConfiguration:         svc.SourceConfiguration.export(),
		AutoScalingConfigurationArn: svc.AutoScalingConfigurationArn,
		HealthCheckConfiguration:    svc.HealthCheckConfiguration.export(),
		InstanceConfiguration:       svc.InstanceConfiguration.export(),
		NetworkConfiguration:        svc.NetworkConfiguration.export(),
		ObservabilityConfiguration:  svc.ObservabilityConfiguration.export(),
	})
	if err != nil {
		return err
	}
	return nil
}

type serviceNotFoundError struct {
	serviceName string
}

func (err *serviceNotFoundError) Error() string {
	return fmt.Sprintf("aarm: service %q is not found", err.serviceName)
}

// getServiceArn finds the service that has name and return its arn.
func (app *App) getServiceArn(ctx context.Context, name string) (string, error) {
	app.Log("[DEBUG] list service to find %q", name)

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
	return "", &serviceNotFoundError{serviceName: name}
}
