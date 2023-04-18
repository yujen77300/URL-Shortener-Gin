package models

import (
	"log"
	// "fmt"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/spf13/viper"
)

var Mongoconn *options.ClientOptions

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	MongodbUrl := viper.GetString("MONGODBURI")

	Mongoconn = options.Client().ApplyURI(MongodbUrl)
	// Mongoclient, err = mongo.Connect(ctx, mongoconn)
	// if err != nil {
	// 	log.Fatal("error while connecting with mongo", err)
	// }
	// err = Mongoclient.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal("error while trying to ping mongo", err)
	// }

	// fmt.Println("mongo connection established")
}
