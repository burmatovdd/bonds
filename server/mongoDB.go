package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func addBondToUserInDB(id string, bondInfo []Bond) {
	var user User
	result := findUserInDBById(id)
	if result.Login != "" {
		fmt.Println("result: ", result)
		for i := 0; i < len(bondInfo); i++ {
			user.Bond = append(user.Bond, Bond{bondInfo[i].Name, bondInfo[i].Count})
		}
		//result.Bond = append(result.Bond, Bond{bondInfo, bondInfo.Count})
		update := bson.D{
			{"$set", bson.D{
				{
					"bond", user.Bond,
				},
			},
			},
		}
		filter := bson.D{{"name", result.Name}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			fmt.Println("update err: ", err)
		}
	}
}

func addBondToDB(bondInfo Bond) {
	res := getBond(bondInfo.Name)
	if res.Name != bondInfo.Name {
		fmt.Println("ready to insert bond")
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

func addUserToDB(user User) bool {
	res := findUserInDB(user.Login)
	if res.Login != user.Login {
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		return true
	} else {
		fmt.Println("user already exist!")
		return false
	}
}

func generateId() primitive.ObjectID {
	id := primitive.NewObjectID()
	var filter bson.D
	var result User
	filter = bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
	}
	if id != result.ID {
		return id
	} else {
		return primitive.NewObjectID()
	}
}

func findUserInDBById(id string) User {
	var result User
	var filter bson.D
	filter = bson.D{{"user_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	return result

}

func findUserInDB(login string) User {
	var result User
	var filter bson.D
	filter = bson.D{{"login", login}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
	}
	return result
}

func getAllUsers() []User {
	findOptions := options.Find()
	var results []User

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	for cur.Next(context.TODO()) {
		var elem User
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
	return results
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
	results := []Bond{}

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
