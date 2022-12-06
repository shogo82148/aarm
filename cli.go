package aarm

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type GlobalOptions struct {
	Region  string
	Profile string
}

func (opts *GlobalOptions) Install(set *flag.FlagSet) {
	set.StringVar(&opts.Region, "region", "", "aws region")
	set.StringVar(&opts.Profile, "profile", "", "aws profile")
}

func (opts *GlobalOptions) newConfig(ctx context.Context) (aws.Config, error) {
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

		cfg, err := opts.newConfig(ctx)
		if err != nil {
			return 1, err
		}
		app := NewApp(cfg)
		if err := app.Init(ctx, &opts); err != nil {
			return 1, err
		}
	case "deploy":
		set := flag.NewFlagSet("aarm", flag.ExitOnError)
		var opts DeployOption
		opts.Install(set)
		set.Parse(args[1:])

		cfg, err := opts.newConfig(ctx)
		if err != nil {
			return 1, err
		}
		app := NewApp(cfg)
		if err := app.Deploy(ctx, &opts); err != nil {
			return 1, err
		}
	default:
		return 1, fmt.Errorf("unknown sub-command: %s", args[0])
	}

	return 0, nil
}
