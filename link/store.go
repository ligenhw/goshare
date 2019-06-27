package link

import (
	"context"
	"time"

	"github.com/ligenhw/goshare/configration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	link *mongo.Collection
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(configration.Conf.MongoDbURI))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}
	link = client.Database("goshare").Collection("link")
}
