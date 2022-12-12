package model

// ParentsDetails 家长情况表
type ParentsDetails struct {
	*Model
	ParentsBaseInfo  []*ParentBaseInfo `json:"parents_base_info"` // 家长基本信息
	SelectBaseInfo   SelectInfo        `json:"select_base_info"`  // 孩子择偶要求
	RecommendContent string            `json:"recommend_content"` // 推荐宣言
	SonNum           int               `json:"son_num"`           // 儿子数
	GirlNum          int               `json:"girl_num"`          // 女子数
}

type ParentBaseInfo struct {
	Age       int    `json:"age"`
	Gender    int    `json:"gender"`     // 性别 0：未定义，1：女，2：男
	Relation  string `json:"relation"`   // 关系类型
	JobStatus int    `json:"job_status"` // 职位状态 0：未定义，1：在职，2：退休
	JobNature int    `json:"job_nature"` // 职位性质 0：未定义，1：私企，2：国企
	Education int    `json:"education"`  // 学历 0：未定义，1：小学；2：初中；3：高中；4：大学本科；5：硕士研究生；6：博士及以上
	Income    int    `json:"income"`     //收入 0：未定义，1：0-2000元；2：2000-5000元；3：5000-1w元；4：1w-2w元；5：2w-3w元；6：3w+
	Extra     string `json:"extra"`      // 其他信息
}

type SelectInfo struct {
	AgeScope       []int  `json:"age_scope"`       // 年龄范围
	EducationScope []int  `json:"education_scope"` // 学历范围
	HeightScope    []int  `json:"height_scope"`    // 身高范围
	IncomeScope    []int  `json:"income_scope"`    // 收入范围
	National       int    `json:"national"`        // 民族
	CityCode       int    `json:"city_code"`       // 居住地
	NativePlace    int    `json:"native_place"`    // 籍贯
	Extra          string `json:"extra"`           // 其他
}
