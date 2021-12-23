package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	gomongo "go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mongo"
)

func InsertMonitorData(monitorData *mongo.MonitorData) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.Timeout)
	defer cancel()

	_, err := mongo.MonitorDataColl.InsertOne(ctx, monitorData)
	if err != nil {
		Logger.Warn("insert one monitor data err", zap.Error(err))
		return err
	}

	return nil
}

func QueryMonitorData(filter bson.M) ([]*mongo.MonitorData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongo.Timeout)
	defer cancel()

	result := make([]*mongo.MonitorData, 0)
	cursor, err := mongo.MonitorDataColl.Find(ctx, filter)
	if err != nil {
		if err == gomongo.ErrNoDocuments {
			return []*mongo.MonitorData{}, nil
		}
		Logger.Warn("query monitor data err", zap.Error(err))
		return nil, err
	}

	if err := cursor.All(ctx, &result); err != nil {
		Logger.Warn("query monitor data err", zap.Error(err))
		return nil, err
	}

	return result, nil
}
