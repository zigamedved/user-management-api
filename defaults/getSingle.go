package defaults

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSingle[T any](collection *mongo.Collection, c *gin.Context, result *T) error {
	ctxTimeout, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	id := c.Param("id")
	documentID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return err1
	}
	filter := bson.M{"_id": documentID}
	options := options.FindOne()

	err2 := collection.FindOne(ctxTimeout, filter, options).Decode(&result)
	if errors.Is(err2, mongo.ErrNoDocuments) {
		return nil
	}
	if err2 != nil {
		return err2
	}

	return nil
}
