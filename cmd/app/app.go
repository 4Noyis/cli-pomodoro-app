package app

import (
	"fmt"

	"github.com/4Noyis/cli-pomodoro-app/utils"
)

var total_pomodoro_time int

func Run() {

	inputValue := utils.InputField("Pomodoro Time: ")
	total_pomodoro_time = total_pomodoro_time + utils.StartTimer(inputValue)
	fmt.Println(total_pomodoro_time)

}
