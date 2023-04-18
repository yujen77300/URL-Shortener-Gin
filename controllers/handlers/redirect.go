package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yujen77300/URL-Shortener-Gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Redirect(c *gin.Context) {
	var shortUrl *models.ShortenURL
	shortenString := c.Param("shortenString")
	var ctx context.Context
	Mongoclient, err = mongo.Connect(ctx, models.Mongoconn)
	defer Mongoclient.Disconnect(ctx)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	urlCollection := Mongoclient.Database("shortenURL").Collection("urls")
	query := bson.D{bson.E{Key: "short_string", Value: &shortenString}}
	urlCollection.FindOne(ctx, query).Decode(&shortUrl)
	if shortUrl == nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"errorMessage": "Shorten URL not found or the original URL is not valid",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, shortUrl.OriginalURL)
}
