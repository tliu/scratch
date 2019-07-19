package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"source.vivint.com/pl/log"
	"source.vivint.com/pl/mongo"
	"time"
)

type test struct {
	Name int    `bson:"a"`
	Data string `bson:"b"`
}

func main() {
	log.Debug("test", log.Fields{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20)*time.Minute)
	defer cancel()
	c := mongo.NewClient().SetAuth("skyPlatform", "mango").SetHost("localhost:27017")
	c = c.SetRetrier(5, .05, 10)

	err := c.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	res := []test{}
	opts := mongo.FindOptions{}
	x := mongo.FieldsToProjection([]string{"b"})
	opts.SetProjection(x)
	err = c.Database("thor").Collection("test").Find(mongo.QueryMap{}, &res, &opts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	var i int64
	i, err = c.Database("thor").Collection("test").Count(bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	/*
		var q interface{}
		q, err = c.Database("thor").Collection("test").InsertOne(bson.M{"a": 10, "b": "foo!!"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(q)
	*/

}
