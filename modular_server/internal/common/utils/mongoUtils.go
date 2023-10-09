package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoConnection(Dsn string, Database string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	config := options.Client().ApplyURI(Dsn)
	client, err := mongo.Connect(ctx, config)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
		return nil, err
	}
	db := client.Database(Database)
	return db, nil
}

func GetObjectsWithFilter[T interface{}](ctx context.Context, collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) ([]*T, error) {
	res, err := collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer res.Close(ctx)
	arr := []*T{}
	for res.Next(ctx) {
		ele := new(T)
		if err := res.Decode(ele); err != nil {
			return nil, err
		}
		arr = append(arr, ele)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return arr, nil
}

func GetOneObectWithFilter[T interface{}](ctx context.Context, collection *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	cur := collection.FindOne(ctx, filter, opts...)
	if err := cur.Err(); err != nil {
		return nil, err
	}
	ele := new(T)
	if err := cur.Decode(ele); err != nil {
		return nil, err
	}
	return ele, nil
}
