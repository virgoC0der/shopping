package models

import (
	"context"
	"shopping/utils/mysql/shopping"

	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func GetUser(username string) (*shopping.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	user := &shopping.User{}
	tx := mysql.GetDB(ctx).Where("username = ?", username).Find(user)
	if tx.Error != nil {
		Logger.Warn("get user err", zap.Error(tx.Error))
		return nil, tx.Error
	}

	return user, nil
}

func GetUserById(userId string) (*shopping.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	user := &shopping.User{}
	tx := mysql.GetDB(ctx).Where("id = ?", userId).Find(user)
	if tx.Error != nil {
		Logger.Warn("get user err", zap.Error(tx.Error))
		return nil, tx.Error
	}

	return user, nil
}

func UpdateUserBalance(userId string, balance int) error {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	tx := mysql.GetDB(ctx).
		Model(&shopping.User{}).
		Where("id = ?", userId).
		Update("balance = balance + ?", balance)
	if tx.Error != nil {
		Logger.Warn("update user err", zap.Error(tx.Error))
		return tx.Error
	}

	return nil
}
