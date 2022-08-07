package main

import (
	"github.com/liwuv587/ffs/cmd/ffs/commands"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := commands.App.Run(os.Args)
	if err != nil {
		logrus.Info(err)
		return
	}
}
