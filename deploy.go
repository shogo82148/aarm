package aarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

func (app *App) Deploy(ctx context.Context) error {
	_, err := app.appRunner.StartDeployment(ctx, &apprunner.StartDeploymentInput{
		ServiceArn: aws.String("arn:aws:apprunner:ap-northeast-1:445285296882:service/my-app-runner-test/c155dbd6ccd44807abae9fef34064205"),
	})
	if err != nil {
		return err
	}

	return nil
}
