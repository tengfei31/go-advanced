package distribute

import (
	"log"
	"testing"
)

func TestShuffle(t *testing.T) {
	var cnt1 = make(map[string]int)
	for i := 0; i < 100000; i++ {
		var sl = []string{
			"127.0.0.1:9000",
			"127.0.0.1:9001",
			"127.0.0.1:9002",
			"127.0.0.1:9003",
			"127.0.0.1:9004",
			"127.0.0.1:9005",
		}
		shuffle1(sl)
		cnt1[sl[0]]++
	}

	var cnt2 = make(map[string]int)
	for i := 0; i < 100000; i++ {
		var sl = []string{
			"127.0.0.1:9000",
			"127.0.0.1:9001",
			"127.0.0.1:9002",
			"127.0.0.1:9003",
			"127.0.0.1:9004",
			"127.0.0.1:9005",
		}
		shuffle2(sl)
		cnt2[sl[0]]++
	}

	log.Println(cnt1)
	log.Println(cnt2)

}
