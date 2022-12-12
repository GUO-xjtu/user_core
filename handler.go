package main

import (
	"context"
	"github.com/GUO-xjtu/kitex_gen/user"
	"github.com/xiangqin/user_core/internal"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// UserRegister implements the UserImpl interface.
func (s *UserImpl) UserRegister(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	resp, err = internal.RegisterNewUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
