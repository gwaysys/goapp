package cmd

import (
	"strings"

	"github.com/urfave/cli/v2"
)

type App struct {
	cli.App
}

func (a *App) Register(category string, cmds ...*cli.Command) {
	if len(cmds) == 0 {
		panic("no command found")
	}
	for _, cmd := range cmds {
		cmd.Category = strings.ToUpper(category)
		a.App.Commands = append(a.App.Commands, cmd)
	}
}
