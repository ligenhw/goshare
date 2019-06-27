package link

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetLinks get all friend links
func GetLinks(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := link.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	var links []Link
	for cur.Next(ctx) {
		var result Link
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(links)
}
