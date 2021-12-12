package services

import (
	"sync"

	"go.uber.org/zap"

	"shopping/apps/api/models"

	. "shopping/utils/log"
)

type ProductCache struct {
	sync.Mutex

	Cache map[int64]int
}

var Product2Count *ProductCache

func InitProductCache() error {
	products, err := models.QueryAllProducts()
	if err != nil {
		Logger.Warn("query all products err", zap.Error(err))
		return err
	}

	if len(products) == 0 {
		return nil
	}

	Product2Count = &ProductCache{
		Cache: make(map[int64]int, len(products)),
	}
	for _, product := range products {
		Product2Count.Lock()
		Product2Count.Cache[product.Id] = product.Amount
		Product2Count.Unlock()
	}

	return nil
}
