package aarm

import (
	"bufio"
	"context"
	"flag"
	"os"
)

type RenderOption struct {
	GlobalOptions
}

func (opts *RenderOption) Install(set *flag.FlagSet) {
	opts.GlobalOptions.Install(set)
}

func (app *App) Render(ctx context.Context, opts *RenderOption) error {
	out := bufio.NewWriter(os.Stdout)
	svc, err := app.loadService(opts.ConfigPath)
	if err != nil {
		return err
	}
	data, err := app.marshalService(svc)
	if err != nil {
		return err
	}
	_, err = out.Write(data)
	if err != nil {
		return err
	}
	return out.Flush()
}
