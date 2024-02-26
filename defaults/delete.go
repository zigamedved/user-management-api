package defaults

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Delete[T any](collection *mongo.Collection, c *gin.Context) error {
	ctxTimeout, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	id := c.Param("id")
	documentID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return err1
	}
	filter := bson.M{"_id": documentID}
	options := options.Delete()

	_, err2 := collection.DeleteOne(ctxTimeout, filter, options)
	if err2 != nil {
		return err2
	}

	return nil
}
