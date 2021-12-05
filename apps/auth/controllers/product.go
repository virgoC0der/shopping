package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/webbase"

	"shopping/apps/auth/io"
	"shopping/apps/auth/models"
)

// GetProduct 获取商品
func GetProduct(c *gin.Context) {
	req := &io.GetProductRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
        c.JSON(200, webbase.ErrInputParams)
        return
    }

	products, err := models.QueryProductById(req.ProductIds)
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		c.JSON(200, webbase.ErrSystemBusy)
        return
	}

	webbase.ServeResponse(c, webbase.ErrOK, products)
}
