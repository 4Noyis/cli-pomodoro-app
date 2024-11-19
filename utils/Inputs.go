package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/4Noyis/cli-pomodoro-app/ui"
)

func InputField(headerText string) int {
	const exitIndicator = -1
	var inputValue int

	for {
		fmt.Print(headerText + "-->")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		input := strings.TrimSpace(line)

		if input == "b" {
			ui.ClearScreen()
			return exitIndicator
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid input or type 'b' to quit")
			continue
		}

		inputValue = value
		break
	}
	return inputValue
}
