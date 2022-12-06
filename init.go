package aarm

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

type InitOption struct{}

func (app *App) Init(ctx context.Context, opt *InitOption) error {
	out, err := app.appRunner.DescribeService(ctx, &apprunner.DescribeServiceInput{
		ServiceArn: aws.String("arn:aws:apprunner:ap-northeast-1:445285296882:service/my-app-runner-test/c155dbd6ccd44807abae9fef34064205"),
	})
	if err != nil {
		return err
	}

	service := importService(out.Service)
	data, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("service.json", data, 0o644)
}
