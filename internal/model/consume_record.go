package model

type ConsumeInfo struct {
	*Model
	Amount      float32 `json:"amount"`       // 消费金额
	ConsumeTime int64   `json:"consume_time"` // 消费时间
	ItemID      int64   `json:"item_id"`      // 消费订单ID
}
