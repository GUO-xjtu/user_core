package main

import (
	"context"
	"github.com/xiangqin/user_core/kitex_gen/user"
	"log"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// UserRegister implements the UserImpl interface.
func (s *UserImpl) UserRegister(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	resp = user.NewRegisterResponse()
	resp.UserID = req.PhoneNum
	log.Println(resp)
	return resp, nil
}
