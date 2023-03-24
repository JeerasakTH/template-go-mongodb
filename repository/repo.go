package routes

import (
	"context"
	"errors"
	"fmt"
	getcollection "template-go-mongodb/collection"
	database "template-go-mongodb/database"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOne(collection string, payload interface{}) (*mongo.InsertOneResult, error) {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.InsertOne(ctx, payload)
	if err != nil {
		return obj, err
	}

	return obj, err
}

func CreateMany(collection string, payload []interface{}) (*mongo.InsertManyResult, error) {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.InsertMany(ctx, payload)
	if err != nil {
		return obj, err
	}

	if obj == nil {
		err := errors.New("INSERT FAIL")
		return obj, err
	}
	return obj, nil
}

func GetOne(collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.FindOne()
	option.SetSort(filterOption)
	defer cancel()

	err := getCollection.FindOne(ctx, filter, option).Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func GetMany(collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.Find()
	option.SetSort(filterOption)
	defer cancel()

	obj, err := getCollection.Find(ctx, filter, option)
	if err != nil {
		return err
	}

	err = obj.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetManyLimit(collection string, filter interface{}, filterOption interface{}, limit int64, data interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.Find()
	option.SetSort(filterOption)
	option.SetLimit(limit)
	defer cancel()

	obj, err := getCollection.Find(ctx, filter, option)
	if err != nil {
		return err
	}

	err = obj.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetManyLimitOne(collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	option := options.Find()
	option.SetSort(filterOption)
	option.SetLimit(1)
	defer cancel()

	obj, err := getCollection.Find(ctx, filter, option)
	if err != nil {
		return err
	}

	err = obj.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(collection string, filter interface{}, set interface{}) (*mongo.UpdateResult, error) {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.UpdateOne(ctx, filter, set)
	if err != nil {
		fmt.Println("obj", obj)
		err := errors.New("update fail, syntax error")
		return obj, err
	}
	if obj == nil {
		fmt.Println("obj", obj)
		err := errors.New("update fail, syntax error")
		return obj, err
	}

	// if obj.MatchedCount == 0 {
	// 	err := errors.New("data mismatched")
	// 	return err
	// }
	// if obj.ModifiedCount == 0 {
	// 	err := errors.New("matched but Nothing changed")
	// 	return err
	// }

	return obj, nil
}

func UpdateMany(collection string, filter interface{}, set interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.UpdateMany(ctx, filter, set)
	if err != nil {
		fmt.Println("obj", obj)
		err := errors.New("update fail, syntax error")
		return err
	}
	if obj == nil {
		fmt.Println("obj", obj)
		err := errors.New("update fail, syntax error")
		return err
	}

	// if obj.MatchedCount == 0 {
	// 	err := errors.New("data mismatched")
	// 	return err
	// }
	// if obj.ModifiedCount == 0 {
	// 	err := errors.New("matched but Nothing changed")
	// 	return err
	// }

	return nil
}

func DeleteOne(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.DeleteOne(ctx, filter)
	if err != nil {
		return obj, err
	}

	if obj.DeletedCount < 1 {
		err := errors.New("no data to delete")
		return obj, err
	}

	return obj, nil
}

// wait test
func DeleteMany(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.DeleteMany(ctx, filter)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

// wait test
func Aggregate(collection string, filter interface{}, data interface{}) error {
	var DB = database.ConnectDB()
	var getCollection = getcollection.GetCollection(DB, collection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	obj, err := getCollection.Aggregate(ctx, filter)
	if err != nil {
		return err
	}

	err = obj.All(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
