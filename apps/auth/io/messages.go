package io

import (
	"shopping/utils/mysql"
)

type LoginRequest struct {
	Username string `json:"username"  binding:"required;username"`
	Password string `json:"password"  binding:"required;password"`
}

type GetProductRequest struct {
	ProductIds []int `form:"product_ids"  binding:"required"`
}

type GetProductResp struct {
	ProductList []*mysql.Product `json:"product_list"`
}
