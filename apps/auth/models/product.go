package models

import (
	"context"

	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func QueryProductById(productId []int) ([]*mysql.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*mysql.Product, 0, len(productId))
    err := mysql.GetDB(ctx).Where("id IN ?", productId).Find(products).Error
    if err != nil {
    	Logger.Warn("query product by id err", zap.Error(err))
        return nil, err
    }

    return products, nil
}
