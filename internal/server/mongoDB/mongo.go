package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	bondsConfig "modules/configs/server"
	"modules/internal/server/UserModel"
	"modules/internal/server/bondsInfo"
)

var collection *mongo.Collection
var ctx = context.TODO()

type MongoService struct {
	method MongoMethods
}

type MongoMethods interface {
	AddBondToUserInDB(id string, bondInfo []bondsInfo.Bond)
	AddUserToDB(user UserModel.User) bool
	GenerateId() primitive.ObjectID
	DeleteBond(name string)
	FindUserInDBById(id string) UserModel.User
	FindUserInDB(login string) UserModel.User
}

// init код запускается после подключения пакета
func init() {
	service := bondsConfig.BondsConfigService{}
	config, err := service.LoadConfig("configs/server")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
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

func find(bonds []bondsInfo.UserBonds, bond bondsInfo.UserBonds) (int, *bondsInfo.UserBonds) {
	for i, item := range bonds {
		if item.Bond.Name == bond.Bond.Name {
			return i, &item
		}
	}
	return -1, nil
}

func (service *MongoService) AddBondToUserInDB(id string, bond bondsInfo.UserBonds) {
	fmt.Println("bondInfo: ", bond)
	result := service.FindUserInDBById(id)
	index, existedBond := find(result.Bonds, bond)

	if existedBond != nil {
		result.Bonds[index].Count = bond.Count
	} else {
		result.Bonds = append(result.Bonds, bondsInfo.UserBonds{Bond: bond.Bond, Count: bond.Count})
	}
	_, err := collection.UpdateOne(ctx,
		bson.M{"user_id": id},
		bson.D{
			{"$set", bson.D{{"bonds", result.Bonds}}},
		})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (service *MongoService) AddUserToDB(user UserModel.User) bool {
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
	var result UserModel.User
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

func (service *MongoService) FindUserInDBById(id string) UserModel.User {
	var result UserModel.User
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
		{"$pull", bson.D{{"bonds", bson.D{{"bond.name", name}}}}},
	}

	//_, err := collection.UpdateMany(ctx, bson.D{}, update)
	_, err := collection.UpdateOne(ctx, bson.D{}, update)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}

}

func (service *MongoService) FindUserInDB(login string) UserModel.User {
	var result UserModel.User
	var filter bson.D
	filter = bson.D{{"login", login}}
	fmt.Println(filter)
	fmt.Println(context.TODO())
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}
