package services

import (
	"sync"

	"go.uber.org/zap"

	"shopping/apps/api/models"

	. "shopping/utils/log"
)

var Product2Count sync.Map

func InitProductCache() error {
	products, err := models.QueryAllProducts()
	if err != nil {
		Logger.Warn("query all products err", zap.Error(err))
		return err
	}

	if len(products) == 0 {
		return nil
	}

	for _, product := range products {
		Product2Count.Store(product.Id, product.Amount)
	}

	return nil
}
