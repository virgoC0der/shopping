package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/auth/io"
	"shopping/apps/auth/models"
)

// GetProductList 获取商品概览
func GetProductList(c *gin.Context) {
	req := &io.GetProductListRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		Logger.Warn("valid get product_list req err",
			zap.Error(err), zap.Any("req", req))
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	count, err := models.QueryProductCount()
	if err != nil {
		Logger.Warn("query product count err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	offset := req.PageSize * req.PageIndex
	products, err := models.QueryProductList(offset, req.PageSize)
	if err != nil {
		Logger.Warn("query product list err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	resp := &io.GetProductListResp{
		Total:       int(count),
		ProductList: make([]*io.ProductInfo, 0, len(products)),
	}

	for _, product := range products {
		p := &io.ProductInfo{
			Id:          product.Id,
			CategoryId:  product.CategoryId,
			Name:        product.Name,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			MainImage:   product.MainImage,
		}
		resp.ProductList = append(resp.ProductList, p)
	}

	webbase.ServeResponse(c, webbase.ErrOK, resp)
}

// GetProduct 获取商品
func GetProduct(c *gin.Context) {
	req := &io.GetProductRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		Logger.Error("valid get product req err",
			zap.Error(err), zap.Any("req", req))
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	products, err := models.QueryProductById(req.ProductIds)
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	webbase.ServeResponse(c, webbase.ErrOK, products)
}
