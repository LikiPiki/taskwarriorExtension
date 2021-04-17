package utilities

import "fmt"

func ProgressBarString(size, completed, all int) string {
	completedPercent := float32(completed) / float32(all) * 100
	progressBar := ""

	progressBar += "["
	for i := 0; i < int(completedPercent/100*float32(size)); i++ {
		progressBar += ColorString("#", "green", "bold")
	}
	for i := 0; i < size-int(completedPercent/100*float32(size)); i++ {
		progressBar += "-"
	}
	progressBar += "]"
	progressBar += fmt.Sprintf(" %.2f%%\n", completedPercent)
	return progressBar
}
