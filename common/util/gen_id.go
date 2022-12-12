package util

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
	// 定义一个全局的 machineID 模拟获取
	// 现实环境中应从 zk 或 etcd 中获取
	machineID uint16
)

// 获取 机器编码ID的 回调函数
func getMachineID() (uint16, error) {
	// machineID 返回nil, 则返回专用IP地址的低16位
	return machineID, nil
}

func InitFlake(machineId uint16) (err error) {
	machineID = machineId
	t, _ := time.Parse("2006-01-02", time.Now().String())
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenUserID 获取全局 ID 的函数
func GenUserID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("需要先初始化以后再执行 GetID 函数 err: %#v \n", err)
		return
	}
	return sonyFlake.NextID()
}
