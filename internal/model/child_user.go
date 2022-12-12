package model

// ChildUser 孩子表
type ChildUser struct {
	*Model
	NickName        string   `json:"nick_name"`
	Age             int      `json:"age"`
	Height          int      `json:"height"`           // 身高 cm
	Weight          int      `json:"weight"`           //  体重 kg
	BirthDay        int64    `json:"birth_day"`        // 出生年月 秒
	CityCode        int8     `json:"city_code"`        // 城市编号
	CityName        string   `json:"city_name"`        // 居住地
	NativePlace     int      `json:"native_place"`     // 籍贯
	ImageUrl        []string `json:"image_url"`        // 照片
	MaritalStatus   int      `json:"marital_status"`   // 婚姻状态 0：未定义；1：未婚；2：离异
	Personality     string   `json:"personality"`      // 性格信息
	WorkName        string   `json:"work_name"`        // 职业
	Relation        string   `json:"relation"`         // 关系类型
	JobStatus       int      `json:"job_status"`       // 职位状态 0：未定义，1：在职，2：退休
	JobNature       int      `json:"job_nature"`       // 职位性质 0：未定义，1：私企，2：国企
	Education       int      `json:"education"`        // 学历 0：未定义，1：小学；2：初中；3：高中；4：大学本科；5：硕士研究生；6：博士及以上
	Income          int      `json:"income"`           //收入 0：未定义，1：0-2000元；2：2000-5000元；3：5000-1w元；4：1w-2w元；5：2w-3w元；6：3w+
	HouseProperty   int      `json:"house_property"`   // 房产信息 0：未定义，1：有房有贷，2：有房无贷，3：无房无贷
	CarProperty     int      `json:"car_property"`     // 车产信息 0：未定义，1：有车有贷，2：有车无贷，3：无车无贷
	SchoolName      string   `json:"school_name"`      // 毕业院校名
	JobReview       int      `json:"job_review"`       // 工作审核状态 0：未审核 1：审核通过 2：审核不通过
	EducationReview int      `json:"education_review"` // 学历审核状态 0：未审核 1：审核通过 2：审核不通过
	ExpectWedTime   int64    `json:"expect_wed_time"`  // 期望结婚时间 秒
}
