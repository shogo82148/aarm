package aarm

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/fatih/color"
	"github.com/fujiwara/logutils"
	"github.com/shogo82148/aarm/internal/apprunneriface"
)

func as[T error](err error) (T, bool) {
	var myErr T
	if errors.As(err, &myErr) {
		return myErr, true
	}
	return myErr, false
}

type appRunner struct {
	apprunneriface.DeploymentStarter
	apprunneriface.ServiceCreator
	apprunneriface.ServiceDescriber
	apprunneriface.ServicesLister
}

type App struct {
	appRunner *appRunner
	logger    *log.Logger
}

func newApp(runner *appRunner, opts *GlobalOptions) *App {
	logger := log.New(io.Discard, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	if opts.Debug {
		logger.SetOutput(newLogFilter(os.Stderr, "DEBUG"))
	} else {
		logger.SetOutput(newLogFilter(os.Stderr, "INFO"))
	}
	return &App{
		appRunner: runner,
		logger:    logger,
	}
}

func newAppRunner(cfg aws.Config) *apprunner.Client {
	return apprunner.NewFromConfig(cfg)
}

func newLogFilter(w io.Writer, minLevel string) *logutils.LevelFilter {
	return &logutils.LevelFilter{
		Levels: []logutils.LogLevel{"DEBUG", "INFO", "WARNING", "ERROR"},
		ModifierFuncs: []logutils.ModifierFunc{
			nil, // DEBUG
			nil, // default
			logutils.Color(color.FgYellow),
			logutils.Color(color.FgRed),
		},
		MinLevel: logutils.LogLevel(minLevel),
		Writer:   w,
	}
}

func (app *App) Log(f string, v ...any) {
	logger := app.logger
	if logger == nil {
		return
	}
	logger.Printf(f, v...)
}
