package context

import (
	"sort"

	"github.com/urfave/cli"
)

type StoolContext struct {
	App     *cli.App
	Verbose bool
}

func NewContext() *StoolContext {
	c := &StoolContext{
		App:     cli.NewApp(),
		Verbose: false,
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := c.App
	app.Name = "stool"
	app.Version = "0.1.0"
	app.Description = "Cord update tool"

	app.Before = c.setupGlobalFlags

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "v",
			Usage: "show more info"},
	}

	return c
}

func (ctx *StoolContext) Run(args []string) error {
	sort.Sort(cli.FlagsByName(ctx.App.Flags))
	sort.Sort(cli.CommandsByName(ctx.App.Commands))

	return ctx.App.Run(args)
}

func (ctx *StoolContext) setupGlobalFlags(c *cli.Context) error {
	ctx.Verbose = c.GlobalBool("v")
	return nil
}
