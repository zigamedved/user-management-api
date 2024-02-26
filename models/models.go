package models

import (
	"context"
	"log"

	db "github.com/zigamedved/user-management-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users *mongo.Collection
var Groups *mongo.Collection

func init() {

	Users = db.GetCollection(db.DB, "users")
	Groups = db.GetCollection(db.DB, "groups")

	_, err1 := Users.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	if err1 != nil {
		log.Printf("failed to create index on Users collection")
		panic(err1)
	}

	_, err2 := Groups.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	})
	if err2 != nil {
		log.Printf("failed to create index on Groups collection")
		panic(err2)
	}

}
