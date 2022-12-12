package dao

import (
	"github.com/GUO-xjtu/kitex_gen/user"
	"github.com/xiangqin/user_core/internal/model"
)

func (s *Service) CreateParentsUser(req *user.RegisterRequest) error {
	info := model.ParentsUser{
		Model:     &model.Model{CreatedOn: req.GetPhoneNum()},
		NickName:  req.GetName(),
		CityCode:  0,
		CityName:  "",
		ImageUrl:  "",
		PhoneNum:  req.GetPhoneNum(),
		WeChatNum: "",
	}
	err := info.Create(s.engine)
	return err
}

func (s Service) UpdateParentsUser() error {
	info := model.ParentsUser{
		Model:     nil,
		NickName:  "",
		CityCode:  0,
		CityName:  "",
		ImageUrl:  "",
		PhoneNum:  0,
		WeChatNum: "",
	}
	err := info.Update(s.engine, info)
	return err
}
