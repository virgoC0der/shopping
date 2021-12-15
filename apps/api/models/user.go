package models

import (
	"context"

	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mysql"
)

func GetUser(username string) (*mysql.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	user := &mysql.User{}
	tx := mysql.GetDB(ctx).Where("username = ?", username).Find(user)
	if tx.Error != nil {
		Logger.Warn("get user err", zap.Error(tx.Error))
		return nil, tx.Error
	}

	return user, nil
}

func GetUserById(userId string) (*mysql.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mysql.Timeout)
	defer cancel()

	user := &mysql.User{}
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

	user := &mysql.User{}
	tx := mysql.GetDB(ctx).Where("id = ?", userId).Find(user)
	if tx.Error != nil {
		Logger.Warn("get user err", zap.Error(tx.Error))
		return tx.Error
	}

	user.Balance = balance
	tx = mysql.GetDB(ctx).Where("id = ?", userId).Update("balance", user.Balance)
	if tx.Error != nil {
		Logger.Warn("update user err", zap.Error(tx.Error))
		return tx.Error
	}

	return nil
}
