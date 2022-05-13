package distribute

import (
	"math/rand"
	"time"
)

//负载均衡计算

func init() {
	rand.Seed(int64(time.Nanosecond))
}

func shuffle1(slice []string) {
	for i := 0; i < len(slice); i++ {
		firstId := rand.Intn(len(slice))
		lastId := rand.Intn(len(slice))
		slice[firstId], slice[lastId] = slice[lastId], slice[firstId]
	}
}

func shuffle2(slice []string) {
	for i := len(slice); i > 0; i-- {
		lastId := i - 1
		idx := rand.Intn(i)
		slice[lastId], slice[idx] = slice[idx], slice[lastId]
	}
}
