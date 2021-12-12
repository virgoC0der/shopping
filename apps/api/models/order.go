package models

import (
	"context"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func InsertOrder(order *mysql.Order) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	tx := mysql.GetDB(ctx).Create(order)
	if tx.Error != nil {
		Logger.Warn("insert order err", zap.Error(tx.Error))
		return 0, tx.Error
	}

	return order.Id, nil
}
