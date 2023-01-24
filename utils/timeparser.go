package utils

import (
	"strconv"
	"time"
)

func ParseTimeToStringRepresentation(arg string) string {
	duration, _ := time.ParseDuration(arg)
	durationInMilli := duration.Milliseconds()
	return strconv.FormatInt(durationInMilli, 10)
}

func ParseTimeToIntSixtyFour(arg string) (int64, error) {
	return strconv.ParseInt(arg, 10, 64)
}

func ValidateTime(duration string, endtime string) bool {
	videoDuration, _ := ParseTimeToIntSixtyFour(duration)
	endTime, _ := ParseTimeToIntSixtyFour(endtime)

	return endTime <= videoDuration
}
