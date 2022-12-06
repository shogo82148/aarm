package aarm

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/shogo82148/aarm/internal/apprunneriface/mock"
)

func TestGetServiceArn(t *testing.T) {
	const pageToken = "new-page"
	const serviceName = "my-app-runner-test"
	const want = "arn:aws:apprunner:ap-northeast-1:445285296882:service/my-app-runner-test/c155dbd6ccd44807abae9fef34064205"
	app := &App{
		appRunner: &appRunner{
			ServicesLister: mock.ListServices(func(ctx context.Context, params *apprunner.ListServicesInput, optFns ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error) {
				if aws.ToString(params.NextToken) == pageToken {
					return &apprunner.ListServicesOutput{
						ServiceSummaryList: []types.ServiceSummary{
							{
								ServiceName: aws.String(serviceName),
								ServiceArn:  aws.String(want),
							},
						},
					}, nil
				}
				return &apprunner.ListServicesOutput{
					ServiceSummaryList: []types.ServiceSummary{
						{
							ServiceName: aws.String("another-service"),
							ServiceArn:  aws.String("arn:aws:apprunner:ap-northeast-1:445285296882:service/another-service/c155dbd6ccd44807abae9fef34064205"),
						},
					},
					NextToken: aws.String(pageToken),
				}, nil
			}),
		},
	}
	got, err := app.getServiceArn(context.Background(), "my-app-runner-test")
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("unexpected result: got %s, want %s", got, want)
	}
}
