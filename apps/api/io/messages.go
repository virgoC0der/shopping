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
	ProductIds []int `form:"product_ids"  binding:"required,max=100"`
}

type GetProductResp struct {
	ProductList []*mysql.Product `json:"product_list"`
}

type PlaceOrderReq struct {
	AddressId int64        `json:"address_id" binding:"required,min=0"`
	Items     []*OrderItem `json:"items"      binding:"required,max=100"`
}

type OrderItem struct {
	ProductId int64   `json:"product_id"  binding:"required,min=1"`
	Price     float64 `json:"price"       binding:"required,min=0"`
	Count     int     `json:"count"       binding:"required,min=1,max=100"`
}

type PlaceOrderResp struct {
	OrderId int64 `json:"order_id"`
}

type TopUpReq struct {
	UserId string `json:"user_id"`
	Money  int    `json:"money"   binding:"required,min=0"`
}

type GetUserInfoResp struct {
	Id       string         `json:"id"`
	Username string         `json:"username"`
	RealName string         `json:"real_name"`
	Phone    string         `json:"phone"`
	Balance  float64        `json:"balance"`
	Orders   []*mysql.Order `json:"orders"`
}
