package distribute

import (
	"log"
	"sync"
	"testing"
)

func TestLockChan(t *testing.T) {
	var lock Lock = NewLock()
	var wg sync.WaitGroup
	var counter int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			// defer lock.UnLock()
			if !lock.Lock() {
				log.Println("lock failed")
				return
			}
			counter++
			log.Print("current counter:", counter)
			lock.UnLock()
		}()
	}
	wg.Wait()
}
