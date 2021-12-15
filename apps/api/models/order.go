package models

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func InsertOrderTrans(order *mysql.Order, productId2Cnt map[int64]int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	err := mysql.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		err = tx.Create(order).Error
		if err != nil {
			return err
		}

		for productId, cnt := range productId2Cnt {
			err = tx.Model(&mysql.Product{}).
				Where("id = ?", productId).
				Update("count", cnt).
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

func QueryOrders(userId string) ([]*mysql.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	orders := make([]*mysql.Order, 0)
	err := mysql.GetDB(ctx).Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		Logger.Warn("query orders err", zap.Error(err))
		return nil, err
	}

	return orders, nil
}
