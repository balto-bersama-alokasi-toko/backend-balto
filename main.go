package main

import (
	database "backend-balto/handler/database"
	"backend-balto/handler/server"
	"backend-balto/handler/usecase/merchant"
	"backend-balto/models"
	"context"
	"sync"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("json")
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	var appConfig models.Configuration
	if err := viper.Unmarshal(&appConfig); err != nil {
		panic(err)
	}
	dbContext := context.Background()
	// Connect Database
	db, err := database.ConnectDb(dbContext, &appConfig.Db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbRepo := database.NewDbRepository(db, appConfig.Dbsecret)
	merchantUsecase := merchant.NewMerchantUsecase(dbRepo)

	serverHttp := server.NewServer(merchantUsecase)
	var wg sync.WaitGroup
	wg.Add(1)
	serverHttp.StartListening(appConfig.Server.Port)
	wg.Wait()
}
