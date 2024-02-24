package utils

import (
	"fmt"
	"time"
)

func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}

func GetTimeNano() int64 {
	return time.Now().UnixNano()
}

func TimestampAddSecond(sec int32) int64 {
	return GetTimestamp() + int64(sec*1000)
}

func FormatSeconds(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%d秒", seconds)
	}

	if seconds < 60*60 {
		minutes := seconds / 60
		seconds = seconds % 60
		if seconds == 0 {
			return fmt.Sprintf("%d分", minutes)
		}
		return fmt.Sprintf("%d分%d秒", minutes, seconds)
	}

	if seconds < 60*60*24 {
		hours := seconds / 60 / 60
		minutes := seconds / 60 % 60
		seconds = seconds % 60
		if minutes == 0 && seconds == 0 {
			return fmt.Sprintf("%d小时", hours)
		} else if seconds == 0 {
			return fmt.Sprintf("%d小时%d分", hours, minutes)
		}
		return fmt.Sprintf("%d小时%d分%d秒", hours, minutes, seconds)
	}

	days := seconds / 60 / 60 / 24
	hours := seconds / 60 / 60 % 24
	minutes := seconds / 60 % 60
	seconds = seconds % 60
	if hours == 0 && minutes == 0 && seconds == 0 {
		return fmt.Sprintf("%d天", days)
	} else if minutes == 0 && seconds == 0 {
		return fmt.Sprintf("%d天%d小时", days, hours)
	} else if seconds == 0 {
		return fmt.Sprintf("%d天%d小时%d分", days, hours, minutes)
	}
	return fmt.Sprintf("%d天%d小时%d分%d秒", days, hours, minutes, seconds)
}

func SecondToNano(v int) time.Duration {
	return time.Second * time.Duration(v)
}

func MinuteToNano(v int) time.Duration {
	return time.Minute * time.Duration(v)
}

func HourToNano(v int) time.Duration {
	return time.Hour * time.Duration(v)
}
