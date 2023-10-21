package utils

import "time"

func String(s string) *string {
	result := &s
	return result
}

func Int(s int) *int {
	result := &s
	return result
}

func UInt(s uint) *uint {
	result := &s
	return result
}

func Time(t time.Time) *time.Time {
	result := &t
	return result
}
