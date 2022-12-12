package internal

import (
	"context"

	"github.com/GUO-xjtu/kitex_gen/user"
	"github.com/xiangqin/user_core/common/util"
	"github.com/xiangqin/user_core/internal/dao"
)

func RegisterNewUser(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = user.NewRegisterResponse()

	userID, genErr := util.GenUserID()
	if genErr != nil {
		return resp, genErr
	}
	svr := dao.NewService(ctx)
	svr.CreateParentsUser(req)
	resp.UserID = int64(userID)

	return resp, nil
}
