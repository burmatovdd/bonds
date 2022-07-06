package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	bondsConfig "modules/internal/server/config"
	"modules/internal/server/structs"
)

var collection *mongo.Collection
var ctx = context.TODO()

type MongoService struct {
	method MongoMethods
}

type MongoMethods interface {
	AddBondToUserInDB(id string, bondInfo []structs.Bond)
	UpdateBond(id string, bondInfo []structs.Bond)
	AddUserToDB(user structs.User) bool
	GenerateId() primitive.ObjectID
	DeleteBond(name string)
	FindUserInDBById(id string) structs.User
	FindUserInDB(login string) structs.User
}

func init() {
	service := bondsConfig.BondsConfigService{}
	config, err := service.LoadConfig("../../configs/server")
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

func (service *MongoService) AddBondToUserInDB(id string, bondInfo []structs.Bond) {
	var user structs.User
	result := service.method.FindUserInDBById(id)
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
			user.Bonds = append(user.Bonds, structs.Bond{Name: bondInfo[i].Name})

		}
	}
	//result.Bond = append(result.Bond, Bond{bondInfo, bondInfo.Count})
	update := bson.D{
		{"$set", bson.D{
			{
				"bond", user.Bonds,
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

func (service *MongoService) UpdateBond(id string, bondInfo []structs.Bond) {
	fmt.Println("Hello! i'm updateBond")
}

func (service *MongoService) AddUserToDB(user structs.User) bool {
	res := service.FindUserInDB(user.Login)
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

func (service *MongoService) GenerateId() primitive.ObjectID {
	id := primitive.NewObjectID()
	var filter bson.D
	var result structs.User
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

func (service *MongoService) FindUserInDBById(id string) structs.User {
	var result structs.User
	var filter bson.D
	filter = bson.D{{"user_id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	return result

}

func (service *MongoService) DeleteBond(name string) {
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

func (service *MongoService) FindUserInDB(login string) structs.User {
	fmt.Println("in findUser")
	var result structs.User
	var filter bson.D
	filter = bson.D{{"login", login}}
	fmt.Println("after filter")
	fmt.Println(filter)
	fmt.Println(context.TODO())
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("after error")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("esli tut to jopa")
	return result
}
