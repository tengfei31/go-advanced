package distribute

import (
	"log"
	"go-advanced/db"
	"time"
)

// redis锁

func incr() {
	var lockKey = "counter_lock"
	var counterKey = "counter"

	//获取锁
	lockRes, err := db.Redis.SetNX(lockKey, 1, 5*time.Second).Result()
	if err != nil || !lockRes {
		log.Printf("%v; lock result:%v", err, lockRes)
		return
	}

	//处理逻辑
	// cntVal, err := db.Redis.Get(counterKey).Int64()
	// if err == nil {
	// 	cntVal++
	// 	_, err = db.Redis.Set(counterKey, cntVal, 0).Result()
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// }

	cntVal, err := db.Redis.Incr(counterKey).Result()
	if err != nil {
		log.Printf("incr failed: %v", err)
	}

	log.Print("current counter is ", cntVal)

	defer func() {
		//释放锁
		unLockRes, err := db.Redis.Del(lockKey).Result()
		if err == nil || unLockRes > 0 {
			log.Print("unlock success")
		} else {
			log.Print("unlock failed:", err)
		}
	} ()
}
