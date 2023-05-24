package utils

import "time"

func ParseDateTime(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	dateTime, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}
