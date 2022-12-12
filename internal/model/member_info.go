package model

type MemberVip struct {
	*Model
	VipType   int   `json:"vip_type"`   // 会员类型
	StartTime int64 `json:"start_time"` // 开始时间
	EndTime   int64 `json:"end_time"`   // 结束时间
	ItemID    int64 `json:"item_id"`    // 会员订单ID
}
