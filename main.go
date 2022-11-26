package main

import (
"github.com/xiangqin/user_core/kitex_gen/user/user"
	"log"
)

func main() {
	svr := user.NewServer(new(UserImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}