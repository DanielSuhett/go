package customer

import (
	"context"
	"log"

	"github.com/DanielSuhett/go/gin/modules/Customer/domain/entities"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

type MongoRepository struct {
	db *mongo.Database
	customer *mongo.Collection
}

func Connect(ctx context.Context) (*MongoRepository, error) {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatal(err)
		return &MongoRepository{}, err
	}

	database :=  client.Database("gin")
	customer := database.Collection("customers")

	return &MongoRepository{db: database, customer: customer}, nil
}

func (mr *MongoRepository) Add(customer entities.Customer) error {
	_, err := mr.customer.InsertOne(ctx, customer)

	return err
}

func (mr *MongoRepository) Get(ID uuid.UUID) (*entities.Customer, error) {
	var customer *entities.Customer

	err := mr.customer.FindOne(ctx, ID).Decode(&customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (mr *MongoRepository) Update(ID uuid.UUID, customer entities.Customer) error {
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
	_, err = mr.customer.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})

	if err != nil {
		return err
	}

	return nil
}
