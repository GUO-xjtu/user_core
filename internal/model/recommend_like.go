package model

// RecommendLike 推荐喜好表
type RecommendLike struct {
	*Model
	AgeScope        []int `json:"age_scope"`         // 年龄范围
	WeightScope     []int `json:"weight_scope"`      // 身高范围
	EducationScope  []int `json:"education_scope"`   // 学历范围
	HeightScope     []int `json:"height_scope"`      // 身高范围
	IncomeScope     []int `json:"income_scope"`      // 收入范围
	National        int   `json:"national"`          // 民族
	CityCode        int   `json:"city_code"`         // 居住地
	NativePlace     int   `json:"native_place"`      // 籍贯
	ChildJobStatus  int   `json:"child_job_status"`  // 职位状态 0：未定义，1：在职，2：退休
	ChildJobNature  int   `json:"child_job_nature"`  // 职位性质 0：未定义，1：私企，2：国企
	HouseProperty   int   `json:"house_property"`    // 房产信息 0：未定义，1：有房有贷，2：有房无贷，3：无房无贷
	CarProperty     int   `json:"car_property"`      // 车产信息 0：未定义，1：有车有贷，2：有车无贷，3：无车无贷
	ParentJobStatus int   `json:"parent_job_status"` // 家长职位状态
}
