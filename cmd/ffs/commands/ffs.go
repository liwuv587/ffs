package commands

import (
	"github.com/liwuv587/ffs/config"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var ffsCommand = cli.Command{
	Name:         "ffs",
	ShortName:    "",
	Aliases:      []string{"run"},
	Usage:        "",
	UsageText:    "",
	Description:  "",
	ArgsUsage:    "",
	Category:     "",
	BashComplete: nil,
	Before: func(cmd *cli.Context) error {
		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
		return nil
	},
	After: nil,
	Action: func(cmd *cli.Context) error {
		zap.L().Debug("service start", zap.Int("port", config.Port))
		return nil
	},
	OnUsageError: nil,
	Subcommands:  nil,
	Flags: []cli.Flag{
		config.FlagPort,
		config.FlagMulticore,
	},
	SkipFlagParsing:        false,
	SkipArgReorder:         false,
	HideHelp:               false,
	Hidden:                 false,
	UseShortOptionHandling: false,
	HelpName:               "",
	CustomHelpTemplate:     "",
}
