package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"

	. "shopping/utils/log"
)

type MongoDBConfig struct {
	host     string `ini:"host"`
	username string `ini:"username"`
	password string `ini:"password"`
	poolSize uint64 `ini:"pool_size"`
	timeout  int    `ini:"timeout"`
	uri      string
}

const (
	kMongoUriTemplate = "mongodb://%s:%s@%s"
	confPath          = "/Users/chensx/Desktop/Go/shopping/utils/mongo/conf/mongo.ini"
)

var (
	client          *mongo.Client
	MonitorDataColl *mongo.Collection
	UserLogColl     *mongo.Collection

	Timeout time.Duration
)

func Init() error {
	conf, err := ini.Load(confPath)
	if err != nil {
		Logger.Warn("load mysql config failed, err", zap.Error(err))
		return err
	}

	opt := &MongoDBConfig{}
	if err = conf.MapTo(opt); err != nil {
		Logger.Warn("map mysql config to struct failed, err", zap.Error(err))
		return err
	}

	Logger.Info("mongo config", zap.Any("opt", opt))
	opt.uri = fmt.Sprintf(kMongoUriTemplate, opt.username, opt.password, opt.host)
	Logger.Info("mongo uri", zap.String("uri", opt.uri))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(opt.uri).SetMaxPoolSize(opt.poolSize))
	if err != nil {
		Logger.Warn("connect mongodb err", zap.Error(err))
		return err
	}

	Timeout = time.Duration(opt.timeout) * time.Second

	InitCollections()
	return nil
}

func InitCollections() {
	MonitorDataColl = client.Database(KDBMonitor).Collection(KCollMonitorData)
	UserLogColl = client.Database(KDBMonitor).Collection(KCollUserLog)
}
