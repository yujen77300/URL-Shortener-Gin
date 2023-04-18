package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yujen77300/URL-Shortener-Gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongoclient *mongo.Client
var err error

func Shorten(c *gin.Context) {
	var ctx context.Context
	Mongoclient, err = mongo.Connect(ctx, models.Mongoconn)
	defer Mongoclient.Disconnect(ctx)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = Mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")
	urlCollection := Mongoclient.Database("shortenURL").Collection("urls")

	document := bson.M{
		"original_url": "https://www.google.com",
		"short_string": "1qaz2wsx",
	}
	ctx = context.TODO()
	result, err := urlCollection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal("error while inserting a document", err)
	}

	fmt.Println("Inserted a document with ID: ", result.InsertedID)

	c.HTML(http.StatusOK, "shorten.html", gin.H{
		"randomstring": "Qtesttse",
	})

}
