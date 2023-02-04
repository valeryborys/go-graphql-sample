package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"vborys/model"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Panic("Mongo connection failed", err)
	}
	collection := client.Database("my-hello-world").Collection("users")
	err = client.Connect(ctx)
	if err != nil {
		log.Panic("Smth wrong with collection", err)
	}
	return collection
}

func FindAllUsers(_ graphql.ResolveParams) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := getCollection()
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print("Error when finding users", err)
		return nil, err
	}
	defer result.Close(ctx)
	var r []model.User
	err = result.All(ctx, &r)
	result.Decode(r)
	if err != nil {
		log.Print("Error when reading users from cursor", err)
	}

	return r, nil
}

func GetById(g graphql.ResolveParams) (interface{}, error) {
	id := g.Args["Id"].(string)
	objId, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := getCollection()
	result := collection.FindOne(ctx, bson.M{"_id": objId})
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		log.Printf("Smth wrong with user %s", id)
	}
	return user, nil
}

func CreateUser(g graphql.ResolveParams) (interface{}, error) {
	user := g.Args["input"]
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := getCollection()
	result, err := collection.InsertOne(ctx, user)
	oi := result.InsertedID.(primitive.ObjectID).Hex()
	objId, _ := primitive.ObjectIDFromHex(oi)
	saved := collection.FindOne(ctx, bson.M{"_id": objId})
	var savedUser model.User
	saved.Decode(&savedUser)
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}

func UpdateUser(g graphql.ResolveParams) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := getCollection()
	id := g.Args["Id"].(string)
	firstName := g.Args["FirstName"].(string)
	lastName := g.Args["LastName"].(string)
	objId, _ := primitive.ObjectIDFromHex(id)
	collection.ReplaceOne(ctx, bson.M{"_id": objId}, bson.M{
		"FirstName": firstName,
		"LastName":  lastName,
	},
	)
	var user model.User
	sr := collection.FindOne(ctx, bson.M{"_id": objId})
	sr.Decode(&user)
	return user, nil
}

func DeleteById(g graphql.ResolveParams) (interface{}, error) {
	id := g.Args["Id"].(string)
	objId, _ := primitive.ObjectIDFromHex(id)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := getCollection()
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objId})
	i := result.DeletedCount
	var message string
	if err != nil {
		message = fmt.Sprintf("Couldn't delete user with id %s", id)
	} else if i == 0 {
		message = fmt.Sprintf("User with id %s was not fonud", id)
	} else {
		message = "Successfully deleted"
	}
	return message, nil
}
