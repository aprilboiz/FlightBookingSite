package pkg

import (
	"time"
)

func ParseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
