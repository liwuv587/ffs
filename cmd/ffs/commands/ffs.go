package commands

import (
	"github.com/liwuv587/ffs/config"
	"github.com/urfave/cli"
	"go.uber.org/zap"
)

var ffsCommand = cli.Command{
	Name:    "ffs",
	Aliases: []string{"run"},
	Flags: []cli.Flag{
		config.FlagPort,
		config.FlagMulticore,
		config.FlagDebug,
	},
	Before: func(cmd *cli.Context) error {
		logger, _ := zap.NewDevelopment()
		zap.ReplaceGlobals(logger)
		return nil
	},
	Action: func(cmd *cli.Context) error {
		zap.L().Debug("service start", zap.Int("port", config.Port))
		return nil
	},
}
