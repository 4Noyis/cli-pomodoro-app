package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/4Noyis/cli-pomodoro-app/utils"
)

type CmdFlags struct {
	WorkTime  int
	BreakTime int
	LoopCount int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	workTimeUsage := "Set the working time in minutes."
	flag.IntVar(&cf.WorkTime, "worktime", 0, workTimeUsage)
	flag.IntVar(&cf.WorkTime, "w", 0, workTimeUsage+" (shorthand)")

	breakTimeUsage := "Set the break time in minutes."
	flag.IntVar(&cf.BreakTime, "breaktime", 0, breakTimeUsage)
	flag.IntVar(&cf.BreakTime, "b", 0, breakTimeUsage+" (shorthand)")

	loopCountUsage := "Set the number of Pomodoro loops. 1 loop = 1 work session + 1 break session."
	flag.IntVar(&cf.LoopCount, "loopcount", 0, loopCountUsage)
	flag.IntVar(&cf.LoopCount, "r", 0, loopCountUsage+" (shorthand)")

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  ./pomodoro -w 30 -b 10 -r 5")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	return &cf
}

func CleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func (cf *CmdFlags) Execute() {
	completedLoops := 0
	var total_pomodoro_time int
	var total_rest_time int
	if cf.LoopCount != 0 {
		for completedLoops < cf.LoopCount {
			if cf.WorkTime != 0 || cf.BreakTime != 0 {
				CleanScreen()
				fmt.Printf("Progress:%sCompleted: %d / %d\n", utils.Repeat(" ", 44), completedLoops, cf.LoopCount)
				fmt.Printf("Time to focus on work...\n")
				total_pomodoro_time = total_pomodoro_time + utils.StartTimer(cf.WorkTime, "=")

				CleanScreen()
				fmt.Printf("Time to rest now...\n")
				total_rest_time = total_rest_time + utils.StartTimer(cf.BreakTime, "~")
				fmt.Println()
				completedLoops++
			} else {
				fmt.Println("Error: All flags (-w, -b, -l) must be provided with non-zero values.")
				flag.Usage()
				os.Exit(1)

			}
		}
	} else {
		fmt.Println("Error: All flags (-w, -b, -l) must be provided with non-zero values.")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", utils.Repeat("=", 50))
	fmt.Printf("%sWell Done!  \n", utils.Repeat(" ", 15))
	fmt.Printf("%s\n", utils.Repeat("=", 50))
	fmt.Printf("%sCompleted Pomodoros:  %d/%d\n", utils.Repeat(" ", 10), completedLoops, cf.LoopCount)
	fmt.Printf("%sTotal Work Time:    %3dm\n", utils.Repeat(" ", 10), total_pomodoro_time)
	fmt.Printf("%sTotal Break Time:   %3dm\n", utils.Repeat(" ", 10), total_rest_time)
	fmt.Printf("%s\n\n", utils.Repeat("=", 50))
}
