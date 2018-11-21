package create

import (
	"fmt"

	"cord.stool/context"
	"cord.stool/updater"
	"github.com/urfave/cli"
	"github.com/kr/pretty"
)

var args = struct {
	SourceDir string
	TargetDir string
	Archive bool
}{
	Archive: false,
}

func Register(ctx *context.StoolContext) {
	cmd := cli.Command{
		Name:        "create",
		Usage:       "Create update",
		Description: "Create update for application",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "sourceDir",
				Usage:       "Path to game(short form sd)",
				Value:       "",
				Destination: &args.SourceDir,
			},
			cli.StringFlag{
				Name:        "sd",
				Hidden:      true,
				Destination: &args.SourceDir,
			},
			cli.StringFlag{
				Name:        "outputDir",
				Usage:       "Path to game(short form sd)",
				Value:       "",
				Destination: &args.TargetDir,
			},
			cli.StringFlag{
				Name:        "od",
				Hidden:      true,
				Destination: &args.TargetDir,
			},
			cli.BoolFlag{
				Name: "archive",
				Destination: &args.Archive,
			},
		},
		Action: func(c *cli.Context) error {
			e := do(ctx, c)

			return e
		},
	}
	ctx.App.Commands = append(ctx.App.Commands, cmd)
}

func do(ctx *context.StoolContext, c *cli.Context) error {
	u, e := updater.PrepairDistr(args.SourceDir, args.TargetDir, false)

	fmt.Printf("%# v", pretty.Formatter(u))


	return e
}