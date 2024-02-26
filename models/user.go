package models

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/zigamedved/user-management-api/database"
	h "github.com/zigamedved/user-management-api/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	h.GetValidator().RegisterValidation("validateGroupID", validateGroupID)
}

type User struct {
	Document `bson:",inline" json:",inline"`
	Email    string             `bson:"email" json:"email" validate:"required,email"`
	Password string             `bson:"password" json:"password" validate:"required"`
	Name     string             `bson:"name" json:"name" validate:"required"`
	GroupId  primitive.ObjectID `bson:"groupID" json:"groupID" validate:"required,validateGroupID"`
}

func validateGroupID(field validator.FieldLevel) bool {
	groupCollection := database.GetCollection(database.DB, "groups")
	groupID := field.Field().Interface().(primitive.ObjectID)

	options := options.Count()
	count, err := groupCollection.CountDocuments(context.Background(), bson.M{"_id": groupID}, options)
	if err != nil {
		return false
	}

	return count == 1
}
