package controllers

import "shopping/utils/webbase"

var (
	ErrPassword = &webbase.CommonResp{
		Code:    2001,
		Message: "用户名或密码错误",
	}

	ErrProductNotEnough = &webbase.CommonResp{
		Code:    2002,
		Message: "商品库存不足",
	}
)
