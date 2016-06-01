package common

import "time"

func GetCurDate() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	return t
}
