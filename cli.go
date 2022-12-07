package aarm

import (
	"context"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type keyValuesOptions struct {
	m map[string]string
}

func (opt keyValuesOptions) String() string {
	keys := make([]string, 0, len(opt.m))
	for k := range opt.m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	kv := make([]string, 0, len(opt.m))
	for _, k := range keys {
		kv = append(kv, k+"="+opt.m[k])
	}
	return strings.Join(kv, ",")
}

func (opt *keyValuesOptions) Set(s string) error {
	if opt.m == nil {
		opt.m = map[string]string{}
	}
	k, v, _ := strings.Cut(s, "=")
	opt.m[k] = v
	return nil
}

type GlobalOptions struct {
	Region     string
	Profile    string
	Debug      bool
	ConfigPath string
	ExtStr     keyValuesOptions
	ExtCode    keyValuesOptions
}

func (opts *GlobalOptions) Install(set *flag.FlagSet) {
	set.StringVar(&opts.Region, "region", "", "aws region")
	set.StringVar(&opts.Profile, "profile", "", "aws profile")
	set.BoolVar(&opts.Debug, "debug", false, "debug")
	set.StringVar(&opts.ConfigPath, "config-path", "aarm.yml", "config path")
	set.Var(&opts.ExtStr, "ext-str", "ext strings")
	set.Var(&opts.ExtCode, "ext-code", "ext code")
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

	case "render":
		set := flag.NewFlagSet("aarm", flag.ExitOnError)
		var opts RenderOption
		opts.Install(set)
		set.Parse(args[1:])

		app, err := NewApp(ctx, &opts.GlobalOptions)
		if err != nil {
			return 1, err
		}
		if err := app.Render(ctx, &opts); err != nil {
			return 1, err
		}

	default:
		return 1, fmt.Errorf("unknown sub-command: %s", args[0])
	}

	return 0, nil
}
