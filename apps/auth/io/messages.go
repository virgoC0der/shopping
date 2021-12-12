package io

import (
	"shopping/utils/mysql"
)

type LoginRequest struct {
	Username string `json:"username"  binding:"required;username"`
	Password string `json:"password"  binding:"required;password"`
}

type GetProductListRequest struct {
	PageIndex int `json:"page_index"  binding:"required;min=0"`
	PageSize  int `json:"page_size"   binding:"required;min=10,max=100"`
}

type GetProductListResp struct {
	Total       int            `json:"total"`
	ProductList []*ProductInfo `json:"product_list"`
}

type ProductInfo struct {
	Id          int64   `json:"id"`
	CategoryId  int64   `json:"category_id"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	MainImage   string  `json:"main_image"`
}

type GetProductRequest struct {
	ProductIds []int `form:"product_ids"  binding:"required"`
}

type GetProductResp struct {
	ProductList []*mysql.Product `json:"product_list"`
}
