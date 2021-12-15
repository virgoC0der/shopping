package models

import (
	"context"
	"gorm.io/gorm"

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

func QueryProductCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	var count int64
	err := mysql.GetDB(ctx).Model(&mysql.Product{}).Count(&count).Error
	if err != nil {
		Logger.Warn("query product count err", zap.Error(err))
		return 0, err
	}

	return count, nil
}

func QueryProductList(offset, limit int) ([]*mysql.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*mysql.Product, 0, limit)
	err := mysql.GetDB(ctx).Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func QueryOneProductById(id int64) (*mysql.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	product := &mysql.Product{}
	err := mysql.GetDB(ctx).Where("id = ?", id).First(product).Error
	if err != nil {
		Logger.Warn("query product by id err", zap.Error(err))
		return nil, err
	}

	return product, nil
}

func QueryAllProducts() ([]*mysql.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	products := make([]*mysql.Product, 0)
	err := mysql.GetDB(ctx).Order("id DESC").Find(&products).Error
	if err != nil {
		Logger.Warn("query all products err", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func UpdateProductsCount(productId2Cnt map[int64]int) error {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	err := mysql.GetDB(ctx).Transaction(func(tx *gorm.DB) error {
		for productId, cnt := range productId2Cnt {
			err := tx.Model(&mysql.Product{}).Where("id = ?", productId).Update("count", cnt).Error
			if err != nil {
				Logger.Warn("update product count err", zap.Error(err))
				return err
			}
		}

		// 返回nil，提交事务
		return nil
	})

	if err != nil {
		Logger.Warn("update products count err", zap.Error(err))
		return err
	}

	return nil
}
