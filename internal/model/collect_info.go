package model

// CollectInfo 收藏表
type CollectInfo struct {
	*Model
	LikeMe []int64 `json:"like_me"` // 喜欢我的用户
	ILike  []int64 `json:"i_like"`  // 我喜欢的用户
}
