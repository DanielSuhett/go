package customer

import (
	"context"

	"github.com/DanielSuhett/go/gin/infra/database"
	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.DB.Collection("customers")
var ctx = context.TODO()

func Add(customer *entities.Customer) error {
	_, err := collection.InsertOne(ctx, customer)

	return err
}

func Get(ID uuid.UUID) (*entities.Customer, error) {
	var customer *entities.Customer

	err := collection.FindOne(ctx, ID).Decode(&customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func Update(ID uuid.UUID, customer *entities.Customer) error {
	c, err := bson.Marshal(customer)

	if err != nil {
		return err
	}

	var update bson.M

	err = bson.Unmarshal(c, &update)

	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: ID}}
	_, err = collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})

	if err != nil {
		return err
	}

	return nil
}
