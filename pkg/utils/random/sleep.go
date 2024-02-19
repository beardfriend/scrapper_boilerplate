package random

import (
	"math/rand"
	"time"
)

func Sleep(min, max float64) {
	rand.Seed(time.Now().UnixMicro())
	randomNumber := min + rand.Float64()*(max-min)
	time.Sleep(time.Millisecond * time.Duration(randomNumber*1000)) // 곱하기 10 또는 적절한 상수로 조절
}
