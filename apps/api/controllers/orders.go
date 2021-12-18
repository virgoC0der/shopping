package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

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

var g singleflight.Group

func PlaceOrder(c *gin.Context) {
	req := &io.PlaceOrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		Logger.Warn("PlaceOrder bind json error", zap.Error(err))
		webbase.ServeResponse(c, webbase.ErrInputParams)
		return
	}

	originCache := make(map[int64]int)
	services.Product2Count.Range(func(key, value interface{}) bool {
		originCache[key.(int64)] = value.(int)
		return true
	})

	var err error
	var price float64
	for _, item := range req.Items {
		count, ok := services.Product2Count.Load(item.ProductId)
		if !ok {
			// 商品不存在，穿透到mysql查询
			count, err, _ = g.Do(strconv.Itoa(int(item.ProductId)), func() (interface{}, error) {
				product, err := models.QueryOneProductById(item.ProductId)
				if err != nil {
					Logger.Warn("PlaceOrder get product error", zap.Error(err))
					return nil, err
				}
				services.Product2Count.Store(item.ProductId, product.Amount)
				return product.Amount, nil
			})
			if err != nil {
				continue
			}
		}

		// 库存不足，不能下单，
		if count.(int) < item.Count {
			Logger.Warn("product not enough", zap.Int64("product_id", item.ProductId))
			webbase.ServeResponse(c, ErrProductNotEnough)
			return
		}

		services.Product2Count.Store(item.ProductId, count.(int)-item.Count)
		price += item.Price * float64(item.Count)
	}

	ctx := webbase.GetUserCtx(c)
	// 检查用户余额是否足够
	if ctx.Balance-price < 0 {
		rollbackCache(originCache, req.Items)
		webbase.ServeResponse(c, ErrBalanceNotEnough)
		return
	}

	items := make([]string, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, strconv.Itoa(int(item.ProductId)))
	}
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

	orderId, err := models.InsertOrderTrans(order, req.Items)
	if err != nil {
		Logger.Warn("insert order err", zap.Error(err))
		rollbackCache(originCache, req.Items)
		webbase.ServeResponse(c, webbase.ErrSystemBusy)
		return
	}

	resp := &io.PlaceOrderResp{
		OrderId: orderId,
	}
	webbase.ServeResponse(c, webbase.ErrOK, resp)
}

// rollbackCache 回滚缓存
func rollbackCache(tempCache map[int64]int, items []*io.OrderItem) {
	for _, item := range items {
		if v, ok := tempCache[item.ProductId]; ok {
			services.Product2Count.Store(item.ProductId, v)
		}
	}
}
