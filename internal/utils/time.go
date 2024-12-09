package utils

import "time"

// FormatTimeString returns the current time in the format "2006-01-02 15:04:05"
func FormatTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
