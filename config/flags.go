package config

import (
	"fmt"
	"github.com/urfave/cli"
)

var (
	// Port 程序运行端口
	Port     int = 9000
	FlagPort     = cli.IntFlag{
		Name:        "port",
		Usage:       fmt.Sprintf("--port %d", Port),
		EnvVar:      "FFS_PORT",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Value:       Port,
		Destination: &Port,
	}

	// Multicore 启用多核
	Multicore     bool = false
	FlagMulticore      = cli.BoolFlag{
		Name:        "multicore",
		Usage:       "--multicore",
		EnvVar:      "FFS_MULTICORE",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: &Multicore,
	}

	// Debug 调试模式
	Debug     bool = false
	FlagDebug      = cli.BoolTFlag{
		Name:        "debug",
		Usage:       "启用 debug 模式",
		EnvVar:      "FFS_DEBUG",
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Destination: &Debug,
	}
)
