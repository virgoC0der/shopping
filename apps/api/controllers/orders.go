package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"shopping/apps/api/io"
	"shopping/apps/api/models"
	"shopping/apps/api/services"
	. "shopping/utils/log"
	"shopping/utils/mysql"
	"shopping/utils/webbase"
)

const (
	kOrderStatus = "待发货"
	kTimeLayout  = "2006-01-02 15:04:05"
)

func PlaceOrder(c *gin.Context) {
	req := &io.PlaceOrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		Logger.Warn("PlaceOrder bind json error", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	var price float64
	originCache := services.Product2Count.Cache
	for _, item := range req.Items {
		services.Product2Count.Lock()
		count, ok := services.Product2Count.Cache[item.ProductId]
		if !ok {
			// 商品不存在，穿透到mysql查询
			product, err := models.QueryOneProductById(item.ProductId)
			if err != nil {
				Logger.Warn("query one product by id err", zap.Error(err),
					zap.Int64("product_id", item.ProductId))
				services.Product2Count.Unlock()
				continue
			}
			services.Product2Count.Cache[item.ProductId] = product.Amount
		}

		// 库存不足，不能下单，
		// todo: 回滚cache
		if count < item.Count {
			services.Product2Count.Unlock()
			services.Product2Count.Cache = originCache
			webbase.ServeResponse(c, ErrProductNotEnough)
			return
		}

		services.Product2Count.Cache[item.ProductId] = count - item.Count
		services.Product2Count.Unlock()
		price += item.Price * float64(item.Count)
	}

	items := make([]string, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, strconv.Itoa(int(item.ProductId)))
	}
	ctx := webbase.GetUserCtx(c)
	nowStr := time.Now().Format(kTimeLayout)
	order := &mysql.Order{
		ProductItem: strings.Join(items, ","),
		TotalPrice:  price,
		UserId:      ctx.UserId,
		Status:      kOrderStatus,
		AddressId:   req.AddressId,
		NickName:    ctx.Nickname,
		Created:     nowStr,
		Updated:     nowStr,
	}

	orderId, err := models.InsertOrder(order)
	if err != nil {
		Logger.Warn("insert order err", zap.Error(err))
		// 回滚cache
		// todo: 回滚product表
		services.Product2Count.Cache = originCache
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	resp := &io.PlaceOrderResp{
		OrderId: orderId,
	}
	webbase.ServeResponse(c, webbase.ErrOK, resp)
}
