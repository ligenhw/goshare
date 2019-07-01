package mongo

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateGetCollection create the http handler
func CreateGetCollection(c *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		cur, err := c.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		var ms []bson.M
		for cur.Next(ctx) {
			var m bson.M
			err := cur.Decode(&m)
			if err != nil {
				log.Fatal(err)
			}
			ms = append(ms, m)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(ms)
	}
}
