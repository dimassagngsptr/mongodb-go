package models

import (
	"errors"
	_ "fmt"
	"log"
	"mongodb-go/src/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  uint64             `bson:"age"`
}

func Find() ([]User, error) {
	db, err := configs.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("users").Find(configs.Ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer csr.Close(configs.Ctx) // pastikan untuk menutup kursor setelah selesai

	result := make([]User, 0)
	for csr.Next(configs.Ctx) {
		var row User
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) == 0 {
		return nil, errors.New("No student data found")
	}

	return result, nil
}

func Insert(item *User) error {
	db, err := configs.Connect()
	if err != nil {
		return err
	}

	result, err := db.Collection("users").InsertOne(configs.Ctx, item)
	if err != nil {
		return err
	}

	item.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func Update(id string, newUser *User) error {
	db, err := configs.Connect()
	if err != nil {
		return err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var selector = bson.M{"_id": objectID}
	var changes = bson.M{"$set": bson.M{"name": newUser.Name, "age": newUser.Age}}
	_, err = db.Collection("users").UpdateOne(configs.Ctx, selector, changes)
	if err != nil {
		return err
	}

	return nil
}
func Delete(id string) error {
	db, err := configs.Connect()
	if err != nil {
		return err
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var selector = bson.M{"_id": objectID}
	_, err = db.Collection("users").DeleteOne(configs.Ctx, selector)
	if err != nil {
		return err
	}
	return nil
}
