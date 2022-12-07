package aarm

import (
	"context"
	"flag"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

type WaitOption struct {
	GlobalOptions
	ServiceArn  string
	ServiceName string
}

func (opts *WaitOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
	set.StringVar(&opts.ServiceArn, "service-arn", "", "service arn")
	set.StringVar(&opts.ServiceName, "service-name", "", "service name")
}

func (app *App) Wait(ctx context.Context, opts *WaitOption) error {
	app.Log("[INFO] start waiting")

	arn, err := app.getServiceArn(ctx, opts.ServiceName)
	if err != nil {
		return err
	}

	for {
		out, err := app.appRunner.ListOperations(ctx, &apprunner.ListOperationsInput{
			ServiceArn: aws.String(arn),
		})
		if err != nil {
			app.Log("[ERROR] failed to list operation: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		if len(out.OperationSummaryList) == 0 {
			time.Sleep(10 * time.Second)
			continue
		}

		// TODO: show log

		if isOperationStable(out.OperationSummaryList[0].Status) {
			break
		}
		time.Sleep(10 * time.Second)
	}

	return nil
}

func isOperationStable(status types.OperationStatus) bool {
	switch status {
	case types.OperationStatusSucceeded, types.OperationStatusFailed,
		types.OperationStatusRollbackSucceeded, types.OperationStatusRollbackFailed:

		return true
	}
	return false
}
