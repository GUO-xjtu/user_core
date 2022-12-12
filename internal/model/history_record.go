package model

type HistoryInfo struct {
	*Model
	BrowseRecord    []BrowseInfo `json:"browse_record"`    // 浏览记录
	RecommendRecord []int64      `json:"recommend_record"` // 推荐记录
}

type BrowseInfo struct {
	UserID     int64 `json:"user_id"`     // 用户ID
	BrowseTime int64 `json:"browse_time"` // 浏览时间 秒
}
