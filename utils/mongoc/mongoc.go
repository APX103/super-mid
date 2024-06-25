package mongoc

import (
	"context"
	"encoding/json"
	"time"

	"apx103.com/super-mid/utils/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientImpl struct {
	conf   *config.BaseConfig
	url    string
	db     string
	Client mongo.Client
}

func (c *MongoClientImpl) GetMongoConnectionPool() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(c.url).
		SetMaxPoolSize(uint64(10))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	c.Client = *client
}

func ParseMongoResult(result bson.M) ([]byte, bool) {
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		logrus.Errorf("failed to marshal to json []byte: %s", err)
		return nil, false
	}
	return jsonData, true
}

func NewMongoClientImpl(conf *config.BaseConfig) *MongoClientImpl {
	c := &MongoClientImpl{conf: conf}
	if conf.MongoConfig.Username == "" && conf.MongoConfig.Password == "" {
		c.url = "mongodb://" + conf.MongoConfig.Url + ":" +
			conf.MongoConfig.Port + "/" + conf.MongoConfig.Database
	} else {
		c.url = "mongodb://" + conf.MongoConfig.Username +
			":" + conf.MongoConfig.Password + "@" +
			conf.MongoConfig.Url + ":" + conf.MongoConfig.Port +
			"/" + conf.MongoConfig.Database
	}
	c.db = conf.MongoConfig.Database
	c.GetMongoConnectionPool()
	return c
}

func (c *MongoClientImpl) Find(collection string, key string, value interface{}) (*[][]byte, bool) {
	coll := c.Client.Database(c.db).Collection(collection)
	var result []bson.M
	q := bson.D{{key, value}}
	if key == "" {
		logrus.Debug("find all")
		logrus.Debugf("collection: %s", collection)
		q = bson.D{}
	}
	cursor, err := coll.Find(context.TODO(), q)
	if err != nil {
		logrus.Infof("failed to get key: %s", err)
		return nil, false
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		logrus.Errorf("Decode error: %s", err)
		return nil, false
	}

	var s [][]byte
	for _, r := range result {
		jsonData, err := json.MarshalIndent(r, "", "    ")
		if err != nil {
			logrus.Errorf("failed to marshal to json []byte: %s", err)
			return nil, false
		}
		s = append(s, jsonData)
	}

	return &s, true
}

func (c *MongoClientImpl) FindOne(collection string, key string, value interface{}) ([]byte, bool) {
	coll := c.Client.Database(c.db).Collection(collection)
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{key, value}}).Decode(&result)
	if err != nil {
		logrus.Infof("failed to get key: %s", err)
		return nil, false
	}
	return ParseMongoResult(result)
}

func (c *MongoClientImpl) UpdateOne(collection string, key string, value interface{}, updateStruct interface{}) bool {
	coll := c.Client.Database(c.db).Collection(collection)
	filter := bson.D{{key, value}}
	_, err := coll.UpdateOne(context.TODO(), filter, bson.M{"$set": updateStruct})
	if err != nil {
		logrus.Infof("failed to update data: %s", err)
		return false
	}
	return true
}

func (c *MongoClientImpl) InsertOne(collection string, data interface{}) bool {
	coll := c.Client.Database(c.db).Collection(collection)
	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		logrus.Infof("failed to insert data: %s", err)
		return false
	}
	return true
}

func (c *MongoClientImpl) DeleteOne(collection string, key string, value interface{}) bool {
	coll := c.Client.Database(c.db).Collection(collection)
	_, err := coll.DeleteOne(context.TODO(), bson.D{{key, value}})
	if err != nil {
		logrus.Infof("failed to delete data: %s", err)
		return false
	}
	return true
}
