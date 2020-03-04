package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	_, err := newMongoCli(os.Args[1])
	if err != nil {
		panic(err)
	}

}

func newMongoCli(addr string) (cli *mongo.Client, err error) {
	addr = fmt.Sprintf("mongodb://%s", addr)

	logrus.Debugf("mongo: %s", addr)

	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
