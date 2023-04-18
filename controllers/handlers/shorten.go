package handlers

import (
	"context"
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
	var shortUrl *models.ShortenURL
	var shortString string
	inputUrl := c.PostForm("inputurl")

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

	urlCollection := Mongoclient.Database("shortenURL").Collection("urls")

	query := bson.D{bson.E{Key: "original_url", Value: &inputUrl}}
	urlCollection.FindOne(ctx, query).Decode(&shortUrl)
	if shortUrl == nil {
		shortString = models.RandomString(6)
		document := bson.M{
			"original_url": inputUrl,
			"short_string": shortString,
		}
		ctx = context.TODO()
		_, err := urlCollection.InsertOne(ctx, document)
		if err != nil {
			log.Fatal("error while inserting a document", err)
		}
	} else {
		shortString = shortUrl.ShortString
	}
	c.HTML(http.StatusOK, "shorten.html", gin.H{
		"randomstring": shortString,
	})

}
