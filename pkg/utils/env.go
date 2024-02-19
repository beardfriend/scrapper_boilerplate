package utils

import "os"

type Env int

const (
	DEV Env = iota + 1
	TEST
)

func GetMode() Env {
	mode := os.Getenv("GO_ENV")
	t := toEnvType(mode)
	return t
}

func toEnvType(envStr string) Env {
	switch envStr {
	case "DEV":
		return DEV
	case "TEST":
		return TEST
	}
	return 0
}
