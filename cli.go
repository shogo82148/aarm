package aarm

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type GlobalOptions struct {
	Region     string
	Profile    string
	Debug      bool
	ConfigPath string
}

func (opts *GlobalOptions) Install(set *flag.FlagSet) {
	set.StringVar(&opts.Region, "region", "", "aws region")
	set.StringVar(&opts.Profile, "profile", "", "aws profile")
	set.BoolVar(&opts.Debug, "debug", false, "debug")
	set.StringVar(&opts.ConfigPath, "config-path", "service.json", "config path")
}

func newAWSConfig(ctx context.Context, opts *GlobalOptions) (aws.Config, error) {
	awsOpts := []func(*config.LoadOptions) error{}
	if opts.Region != "" {
		awsOpts = append(awsOpts, config.WithRegion(opts.Region))
	}
	if opts.Profile != "" {
		awsOpts = append(awsOpts, config.WithSharedConfigProfile(opts.Profile))
	}
	return config.LoadDefaultConfig(ctx, awsOpts...)
}

func CLI(ctx context.Context, args []string) (int, error) {
	if len(args) == 0 {
		args = []string{"help"}
	}
	switch args[0] {
	case "init":
		set := flag.NewFlagSet("aarm", flag.ExitOnError)
		var opts InitOption
		opts.Install(set)
		set.Parse(args[1:])

		app, err := NewApp(ctx, &opts.GlobalOptions)
		if err != nil {
			return 1, err
		}
		if err := app.Init(ctx, &opts); err != nil {
			return 1, err
		}

	case "deploy":
		set := flag.NewFlagSet("aarm", flag.ExitOnError)
		var opts DeployOption
		opts.Install(set)
		set.Parse(args[1:])

		app, err := NewApp(ctx, &opts.GlobalOptions)
		if err != nil {
			return 1, err
		}
		if err := app.Deploy(ctx, &opts); err != nil {
			return 1, err
		}

	case "diff":
		set := flag.NewFlagSet("aarm", flag.ExitOnError)
		var opts DiffOption
		opts.Install(set)
		set.Parse(args[1:])

		app, err := NewApp(ctx, &opts.GlobalOptions)
		if err != nil {
			return 1, err
		}
		if err := app.Diff(ctx, &opts); err != nil {
			return 1, err
		}
	default:
		return 1, fmt.Errorf("unknown sub-command: %s", args[0])
	}

	return 0, nil
}
