package format

import (
	"fmt"
	"time"
)

func DateTimeToUnix(str string) int64 {
	time, err := time.Parse(time.RFC3339, str)
	if err != nil {
		fmt.Printf("Error parsing arrive date: %v\n", err)
	}
	return time.Unix()
}
