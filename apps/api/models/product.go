package models

import (
	"context"
	"go.uber.org/zap"
	"shopping/utils/mysql/shopping"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func QueryProductById(productId []int) ([]*shopping.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*shopping.Product, 0, len(productId))
	err := mysql.GetDB(ctx).Where("id IN ?", productId).Find(&products).Error
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func QueryProductCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	var count int64
	err := mysql.GetDB(ctx).Model(&shopping.Product{}).Count(&count).Error
	if err != nil {
		Logger.Warn("query product count err", zap.Error(err))
		return 0, err
	}

	return count, nil
}

func QueryProductList(offset, limit int) ([]*shopping.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*shopping.Product, 0, limit)
	err := mysql.GetDB(ctx).Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func QueryOneProductById(id int64) (*shopping.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	product := &shopping.Product{}
	err := mysql.GetDB(ctx).Where("id = ?", id).First(product).Error
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		return nil, err
	}

	return product, nil
}

func QueryAllProducts() ([]*shopping.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*shopping.Product, 0)
	err := mysql.GetDB(ctx).Order("id DESC").Find(&products).Error
	if err != nil {
		Logger.Warn("query all products err", zap.Error(err))
		return nil, err
	}

	return products, nil
}
