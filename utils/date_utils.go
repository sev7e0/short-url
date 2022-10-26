package utils

import (
	"fmt"
	"time"
)

func GetNowFormatTodayTime() string {
	now := time.Now()
	return fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())
}
