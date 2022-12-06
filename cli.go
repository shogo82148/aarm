package aarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
)

func CLI(ctx context.Context, args []string) (int, error) {
	if len(args) == 0 {
		args = []string{"help"}
	}
	switch args[0] {
	case "init":
		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return 1, err
		}
		app := NewApp(cfg)
		app.Init(ctx, nil)
	}

	return 0, nil
}
