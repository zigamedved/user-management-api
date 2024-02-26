package defaults

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll[T any](collection *mongo.Collection, c *gin.Context, results *[]T) error {
	ctxTimeout, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	options := options.Find()

	cursor, err1 := collection.Find(ctxTimeout, bson.M{}, options)
	if err1 != nil {
		return err1
	}

	if err2 := cursor.All(ctxTimeout, results); err2 != nil {
		return err2
	}

	return nil
}
