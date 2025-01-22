package app

import (
	"github.com/4Noyis/cli-pomodoro-app/commands"
)

func Run() {

	cmdFlags := commands.NewCmdFlags()
	cmdFlags.Execute()

}
