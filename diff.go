package aarm

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/fatih/color"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

type DiffOption struct {
	GlobalOptions
}

func (opts *DiffOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
}

func (app *App) Diff(ctx context.Context, opts *DiffOption) error {
	// local service definition
	configPath := opts.ConfigPath
	svc, err := app.loadService(configPath)
	if err != nil {
		return err
	}
	local, err := app.marshalService(svc)
	if err != nil {
		return err
	}

	// remote service definition
	arn, err := app.getServiceArn(ctx, aws.ToString(svc.ServiceName))
	if err != nil {
		return err
	}
	out, err := app.appRunner.DescribeService(ctx, &apprunner.DescribeServiceInput{
		ServiceArn: aws.String(arn),
	})
	if err != nil {
		return err
	}
	svc = importService(out.Service)
	remote, err := app.marshalService(svc)
	if err != nil {
		return err
	}

	edits := myers.ComputeEdits(span.URIFromPath(arn), string(remote), string(local))
	unified := fmt.Sprintf("%s", gotextdiff.ToUnified(arn, configPath, string(remote), edits))
	fmt.Print(coloredDiff(unified))

	return nil
}

func coloredDiff(src string) string {
	var b strings.Builder
	for _, line := range strings.Split(src, "\n") {
		if strings.HasPrefix(line, "-") {
			b.WriteString(color.RedString(line) + "\n")
		} else if strings.HasPrefix(line, "+") {
			b.WriteString(color.GreenString(line) + "\n")
		} else {
			b.WriteString(line + "\n")
		}
	}
	return b.String()
}
