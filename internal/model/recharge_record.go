package model

// RechargeInfo 充值记录表
type RechargeInfo struct {
	*Model
	Amount       float32 `json:"amount"`        // 充值金额
	RechargeTime int64   `json:"recharge_time"` // 充值时间
}
