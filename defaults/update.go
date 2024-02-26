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

func Update[T any](collection *mongo.Collection, c *gin.Context, document T) error {
	ctxTimeout, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	id := c.Param("id") // could implement GetID method on document to get the ID
	documentID, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		return err1
	}

	filter := bson.M{"_id": documentID}
	update := bson.M{
		"$set": struct {
			ID        primitive.ObjectID `bson:"_id"`
			Document  T                  `bson:",inline"`
			UpdatedAt time.Time          `bson:"updatedAt"`
		}{
			ID:        documentID,
			Document:  document,
			UpdatedAt: time.Now(),
		},
	}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err2 := collection.FindOneAndUpdate(ctxTimeout, filter, update, options).Decode(&document)
	if err2 != nil {
		return err2
	}

	return nil
}
