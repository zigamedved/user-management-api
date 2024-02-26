package defaults

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create[T any](collection *mongo.Collection, c *gin.Context, document T) (*primitive.ObjectID, error) {
	ctxTimeout, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	create := struct {
		ID        primitive.ObjectID `bson:"_id" json:"_id"`
		Document  T                  `bson:",inline"`
		UpdatedAt time.Time          `bson:"updatedAt"`
		CreatedAt time.Time          `bson:"createdAt"`
	}{
		ID:        primitive.NewObjectID(),
		Document:  document,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	res, err := collection.InsertOne(ctxTimeout, create)
	if err != nil {
		return nil, err
	}

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, fmt.Errorf("error while retrieving ID of inserted object")
	}
	return &insertedID, nil
}
