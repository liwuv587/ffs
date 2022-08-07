package commands

import (
	"github.com/urfave/cli"
)

var App *cli.App

func init() {

	App = cli.NewApp()

	// 子命令参数绑定
	App.Flags = append(App.Flags, ffsCommand.Flags...)

	// 设置子命令
	App.Commands = append(App.Commands, ffsCommand)

	// 设置默认运行命令
	App.Action = ffsCommand.Action
	App.Before = ffsCommand.Before
}
