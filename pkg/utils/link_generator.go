package utils

import (
	"math/rand"
	"time"
)

func LinkGenerate() int64 {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int63()
	return n
}
