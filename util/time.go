package util

import (
	"time"
)

func GetZeroDayTime(t time.Time) (dayStartAt time.Time, dayEndAt time.Time) {
	year, month, day := t.Date()

	startAt := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	endAt := startAt.Add(time.Hour * 24)

	return startAt, endAt
}

func GetZeroDayTimeSec(sec int64) (dayStartAtSec, dayEndAtSec int64) {
	t := time.Unix(sec, 0)
	year, month, day := t.Date()

	startAt := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	endAt := startAt.Add(time.Hour * 24)

	return startAt.Unix(), endAt.Unix()
}

//获取相差时间(单位秒)
func GetSecondDiffer(startTime, endTime time.Time) int64 {
	var diff int64

	if endTime.Before(startTime) {
		diff := startTime.Unix() - endTime.Unix()
		return diff
	} else {
		return diff
	}
}
