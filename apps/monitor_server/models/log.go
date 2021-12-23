package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mongo"
)

func InsertUserLog(log *mongo.UserLog) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.Timeout)
	defer cancel()

	_, err := mongo.UserLogColl.InsertOne(ctx, log)
	if err != nil {
		Logger.Warn("insert user log err", zap.Error(err))
		return err
	}

	return err
}

func QueryUserLog(limit, offset int64) (int64, []*mongo.UserLog, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.Timeout)
	defer cancel()

	total, err := mongo.UserLogColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		Logger.Warn("query user log count err", zap.Error(err))
		return 0, nil, err
	}

	opt := options.Find().SetSkip(offset).SetLimit(limit)
	cur, err := mongo.UserLogColl.Find(ctx, nil, opt)
	if err != nil {
		Logger.Warn("query user log err", zap.Error(err))
		return 0, nil, err
	}

	var logs []*mongo.UserLog
	for cur.Next(ctx) {
		var log mongo.UserLog
		err := cur.Decode(&log)
		if err != nil {
			Logger.Warn("query user log err", zap.Error(err))
			continue
		}
		logs = append(logs, &log)
	}

	return total, logs, nil
}
