package distribute

import (
	"fmt"
	"math"
	"time"

	"go-advanced/db"

	"github.com/sony/sonyflake"
)

//分布式ID

func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	var tmpMachineID uint64
	// 节点ID可以是在某个文件里，或者在缓存里面，需要提前放置
	tmpMachineID, err = db.Redis.Get("machine_id").Uint64()
	if err != nil {
		return machineID, err
	}
	if tmpMachineID > math.MaxUint16 {
		return machineID, fmt.Errorf("节点ID超出uint16限制")
	}
	machineID = uint16(tmpMachineID)
	return machineID, err
}

func checkMachineID(machineID uint16) bool {
	result, err := db.Redis.SAdd("machine_ids", machineID).Result()
	if err == nil || result == 1 {
		return true
	}
	_, err = db.Redis.Set("machine_id", result, 0).Result()
	return err == nil
}

func generateSonyflake() (uint64, error) {
	t, _ := time.Parse("2006-01-02", "2022-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)
	if sf == nil {
		return 0, fmt.Errorf("生成失败")
	}
	return sf.NextID()
}
