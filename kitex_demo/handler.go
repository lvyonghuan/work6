package main

import (
	"context"
	"strconv"
	api "work6/kitex_demo/kitex_gen/api"
)

// PickImpl implements the last service interface defined in the IDL.
type PickImpl struct{}

// Pick implements the PickImpl interface.
func (s *PickImpl) Pick(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	var str string
	if req.NickName != nil {
		str = " 序号：" + strconv.FormatInt(req.Count, 10) + " 姓名：" + req.Name + " 年龄:" + strconv.FormatInt(int64(req.Age), 10) + " 昵称：" + *req.NickName
	} else {
		str = " 序号：" + strconv.FormatInt(req.Count, 10) + " 姓名：" + req.Name + " 年龄：" + strconv.FormatInt(int64(req.Age), 10)
	}
	return &api.Response{Indentity: str}, nil
}
