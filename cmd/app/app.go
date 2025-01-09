package app

import (
	"github.com/4Noyis/cli-pomodoro-app/commands"
)

var total_pomodoro_time int
var total_rest_time int

func Run() {

	cmdFlags := commands.NewCmdFlags()
	cmdFlags.Execute()

	/* Getinput from user with CLI then start pomodoro */

	// pomodoroTime := utils.InputField("Pomodoro Time: ")
	// restTime := utils.InputField("Rest Time: ")

	// total_pomodoro_time = total_pomodoro_time + utils.StartTimer(pomodoroTime)
	// println("Good Work! Time to rest now...")

	// total_rest_time = total_rest_time + utils.StartTimer(restTime)

	// fmt.Println(total_rest_time)     // second
	// fmt.Println(total_pomodoro_time) // second

}
