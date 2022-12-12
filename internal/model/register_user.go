package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ParentsUser 注册用户表
type ParentsUser struct {
	*Model
	NickName       string  `json:"nick_name"`       // 昵称
	CityCode       int8    `json:"city_code"`       // 城市编号
	CityName       string  `json:"city_name"`       // 城市名
	ImageUrl       string  `json:"image_url"`       // 头像
	CardNum        int64   `json:"card_num"`        // 身份证号
	PhoneNum       int64   `json:"phone_num"`       // 电话
	WeChatNum      string  `json:"we_chat_num"`     // 微信号
	PassWord       string  `json:"pass_word"`       // 密码
	GrowthValue    int8    `json:"growth_value"`    // 成长值
	AccountBalance float32 `json:"account_balance"` // 账户余额
}

// UserInfo 用户信息
type UserInfo struct {
	Age           int       `json:"age"`
	Gender        int       `json:"gender"` // 性别 0：未定义，1：女，2：男
	Birthday      time.Time `json:"birthday"`
	Relation      string    `json:"relation"`       // 关系类型
	JobStatus     int       `json:"job_status"`     // 职位状态 0：未定义，1：在职，2：退休
	JobNature     int       `json:"job_nature"`     // 职位性质 0：未定义，1：私企，2：国企
	HouseProperty int       `json:"house_property"` // 房产信息 0：未定义，1：有房有贷，2：有房无贷，3：无房无贷
	CarProperty   int       `json:"car_property"`   // 车产信息 0：未定义，1：有车有贷，2：有车无贷，3：无车无贷
	Extra         string    `json:"extra"`          // 其他信息
}

// Create 新增注册用户
func (ru *ParentsUser) Create(db *gorm.DB) error {
	if err := db.Create(&ru).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新注册用户信息
func (ru *ParentsUser) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&ru).Where("ID = ? AND is_del = ?", ru.ID, 0).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

// Delete 删除注册用户信息
func (ru *ParentsUser) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", ru.ID, 0).Delete(&ru).Error; err != nil {
		return err
	}
	return nil
}

// ConditionsQuery 条件查询
func (ru *ParentsUser) ConditionsQuery(db *gorm.DB, oddsMap map[string][]int64, pageSize, pageOffset int) ([]*ParentsUser, error) {
	var ParentsUsers []*ParentsUser

	if ru.ID > 0 {
		if err := db.Where("ID = ? AND is_del = ?", ru.ID, 0).Find(ParentsUsers).Error; err != nil {
			return nil, err
		}
		return ParentsUsers, nil
	}
	if ru.CityCode > 0 {
		db = db.Where("city_code = ?", ru.CityCode)
	}

	for key, fields := range oddsMap {
		switch key {
		case "job_status":
			db = db.Where("self_user_info.job_status IN ?", fields)
		case "job_nature":
			db = db.Where("self_user_info.job_nature IN ?", fields)
		case "house_property":
			db = db.Where("self_user_info.house_property IN ?", fields)
		case "car_property":
			db = db.Where("self_user_info.car_property IN ?", fields)
		}
	}
	if pageOffset > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset(pageOffset * pageSize)
	}
	if err := db.Find(ParentsUsers).Error; err != nil {
		return nil, err
	}
	return ParentsUsers, nil
}
