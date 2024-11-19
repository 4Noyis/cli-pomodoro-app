package utils

import (
	"fmt"
	"time"
)

func StartTimer(pomodoroTime int) int {
	totalDuration := time.Duration(pomodoroTime) * time.Minute // Total duration
	totalSeconds := int(totalDuration.Seconds())
	barWidth := 50 // Width of the progress bar

	fmt.Println("Progress:")

	for elapsed := 0; elapsed <= totalSeconds; elapsed++ {
		// Calculate the remaining time
		remainingTime := totalSeconds - elapsed
		minutes := remainingTime / 60
		seconds := remainingTime % 60

		// Calculate the percentage of completion
		percent := float64(elapsed) / float64(totalSeconds)

		// Calculate the number of '=' to fill in the progress bar
		filledWidth := int(percent * float64(barWidth))
		bar := fmt.Sprintf("[%s%s] %02d:%02d remaining",
			repeat("=", filledWidth),          // Filled part
			repeat(" ", barWidth-filledWidth), // Empty part
			minutes,                           // Minutes remaining
			seconds,                           // Seconds remaining
		)

		// Print the progress bar
		fmt.Printf("\r%s", bar)

		// Sleep for 1 second before updating the progress
		time.Sleep(1 * time.Second)
	}

	// Final output to indicate completion
	fmt.Println("\nDone!")
	return totalSeconds
}

// Helper function to repeat a character `n` times
func repeat(char string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}