package format

import (
	"fmt"
	"ota/constant"
	"strings"
	"time"
)

func MinutesToFormatted(minutes int) string {
	hours := minutes / constant.Minute
	mins := minutes % constant.Minute
	formatted := ""
	if hours > 0 {
		formatted += fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 {
		//if there is hour exist, add space before add minutes
		if formatted != "" {
			formatted += " "
		}
		formatted += fmt.Sprintf("%dm", mins)
	}

	return formatted
}

func FormattedToMinutes(formatted string) int {
	formatted = strings.ReplaceAll(formatted, " ", "")
	duration, err := time.ParseDuration(formatted)
	if err != nil {
		fmt.Printf("Failed to convert formatted to minutes, formatted: %s, err: %v", formatted, err)
	}
	mins := int(duration.Minutes())
	return mins
}
