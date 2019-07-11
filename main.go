package main

import (
	"fmt"
	"source.vivint.com/pl/mongo"
	"source.vivint.com/pl/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*mongoT    imeout)*time.Minute)
	defer cancel()
	c := mongo.NewClient().SetAuth("skyPlatform", "mango").SetHost("mongodb://localhost:27017")
	err := c.Connect(ctx)


}
