package utils

import (
	"strconv"
	"time"
)

func PtrInt(v string) int {

	//v1, _ := strconv.Atoi(v)
	v1, _ := strconv.ParseFloat(v, 64)
	return int(v1)
}

func PtrTime(v string) time.Time {
	t, _ := time.ParseInLocation(time.DateTime, v, time.Local)
	return t.Local()
}
