package dataAccessLayer

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type DailyModel struct {
	Hours   int       `bson:"hours"`
	Project string    `bson:"project"`
	Time    time.Time `bson:"time"`
}

func getClient() (mongo.Client, context.Context) {
	ctx := context.Background()
	client, _ := mongo.Connect(ctx, "mongodb://127.0.0.1:27017")
	return *client, ctx
}

func AddHours(project string, hours int) {
	client, ctx := getClient()
	collection := client.Database("podemos-aprender").Collection("daily")

	newObj := bson.D{
		{"hours", hours},
		{"project", project},
		{"time", time.Now()},
	}

	collection.InsertOne(ctx, newObj)
}

func GetTimeInvested(project string) int {
	client, ctx := getClient()
	collection := client.Database("podemos-aprender").Collection("daily")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"project": project},
		},
		bson.M{
			"$group": bson.M{"_id": "project", "hours": bson.M{"$sum": "$hours"}},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return -1
	}
	defer cursor.Close(context.Background())

	type sumModel struct {
		Hours int `bson:"hours"`
	}
	itemRead := sumModel{}
	cursor.Next(context.Background())
	cursor.Decode(&itemRead)

	return itemRead.Hours

}
