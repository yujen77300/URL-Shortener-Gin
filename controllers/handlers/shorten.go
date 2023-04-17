package handlers

import (
	"context"
	"fmt"
	"log"
	"github.com/yujen77300/URL-Shortener-Gin/models"
	"github.com/gin-gonic/gin"
		"go.mongodb.org/mongo-driver/bson"
)

func Shorten(c *gin.Context){
	var ctx context.Context
	urlCollection := models.Mongoclient.Database("shortenURL").Collection("urls")

	document := bson.M{
		"url": "https://example.com",
		"short": "test",
	}
	ctx = context.TODO()
	result, err := urlCollection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal("error while inserting a document", err)
	}

	fmt.Println("Inserted a document with ID: ", result.InsertedID)

}