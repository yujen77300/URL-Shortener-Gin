package models

import (
	"log"
	"go.mongodb.org/mongo-driver/mongo/options"
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
}
