package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	//value := viper.Get("MONGODB_CONNSTRING")
	//if value == "" {
	//	value = "mongodb://localhost:27017"
	//}
	// из docker-compose.yml  подставляем переменную
	clientOptions := options.Client().ApplyURI(config.MONGOCONN) // problem?
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
			//var res struct {
			//	name  string
			//	count int
			//}
			//err := collection.FindOne(
			//	ctx,
			//	bson.D{
			//		{"bond", bson.D{
			//			{"name", "RU000A1049P5"},
			//		}},
			//	}).Decode(&res)
			//fmt.Println("err:", err)
			user.Bond = append(user.Bond, Bond{bondInfo[i].Name, bondInfo[i].Count})

		}
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

func updateBond(id string, bondInfo []Bond) {
	fmt.Println("Hello! i'm updateBond")
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

func deleteBond(name string) {
	update := bson.D{
		{"$pull", bson.D{
			{
				"bond", bson.D{
					{
						"name", name,
					},
				},
			},
		},
		},
	}

	_, err := collection.UpdateMany(ctx, bson.D{}, update)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

}
