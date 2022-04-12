package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	value := os.Getenv("MONGODB_CONNSTRING")
	if value == "" {
		value = "mongodb://localhost:27017"
	}
	// из docker-compose.yml  подставляем переменную
	clientOptions := options.Client().ApplyURI(value)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database("BondsDB").Collection("bonds")

}

func addToDB(bondInfo Bond) {
	res := getBond(bondInfo.Name)
	if res.Name != bondInfo.Name {
		_, err := collection.InsertOne(ctx, bondInfo)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		if res.Count != bondInfo.Count {
			// update count
			update := bson.D{
				{"$set", bson.D{
					{"count", bondInfo.Count},
				},
				},
			}
			filter := bson.D{{"name", res.Name}}
			_, err := collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				fmt.Println("update err: ", err)
			}
		}
	}
}

func getBond(name string) Bond {
	var result Bond

	filter := bson.D{{"name", name}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
	}
	return result
}

func getAllBonds() []Bond {
	findOptions := options.Find()
	var results []Bond

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	for cur.Next(context.TODO()) {
		var elem Bond
		err := cur.Decode(&elem)
		results = append(results, elem)
		if err != nil {
			fmt.Println("cur.Next err: ", err)
		}

	}

	if err = cur.Err(); err != nil {
		fmt.Println("cur.Err: ", err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())
	fmt.Println(results)
	return results
}

func deleteBond(name string) {
	_, err := collection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		fmt.Println("err: ", err)
	}
	//deleteAllBonds()
}

func deleteAllBonds() {
	_, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
}
