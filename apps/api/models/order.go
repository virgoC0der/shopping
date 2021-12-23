package models

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"shopping/apps/api/io"
	"shopping/utils/mysql/shopping"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func InsertOrderTrans(order *shopping.Order, items []*io.OrderItem) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	err := mysql.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		// 更新用户余额
		err = tx.Model(&shopping.User{}).
			Where("id = ?", order.UserId).
			Update("balance = balance - ?", order.TotalPrice).Error
		if err != nil {
			Logger.Warn("update user balance err", zap.Error(err))
			return err
		}

		// 插入订单
		err = tx.Create(order).Error
		if err != nil {
			Logger.Warn("insert order err", zap.Error(err))
			return err
		}

		// 更新商品数量
		for _, item := range items {
			err = tx.Model(&shopping.Product{}).
				Where("id = ?", item.ProductId).
				Update("count = count - ?", item.Count).
				Error
			if err != nil {
				Logger.Warn("update product count err", zap.Error(err))
				return err
			}
		}

		return nil
	})

	if err != nil {
		Logger.Warn("insert order err", zap.Error(err))
		return 0, err
	}

	return order.Id, nil
}

func QueryOrders(userId string) ([]*shopping.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	orders := make([]*shopping.Order, 0)
	err := mysql.GetDB(ctx).Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		Logger.Warn("query orders err", zap.Error(err))
		return nil, err
	}

	return orders, nil
}
