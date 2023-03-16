package main

import (
	"blog-server-app/DB"
	middleware "blog-server-app/modules/system/middlewares"
	router "blog-server-app/routes"
	"fmt"
	"log"
	"net/http"

	appLogger "blog-server-app/modules/system/services"

	config "github.com/spf13/viper"
)

func main() {

	config.SetConfigFile("config/default.json")

	err := config.ReadInConfig()

	if err != nil {
		log.Fatalln("Failed to read config")
	}

	logger := appLogger.NewAppLogger()

	port := config.GetString("port")

	if port == "" {
		log.Println("No PORT set. Setting it to default 3000")
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)

	//Initiate db connection
	db := DB.New(logger.Named("main/db")).InitConnection()

	//Initialize routes
	router := router.NewRouter(db, logger.Named("main/router"))

	wrappedMux := middleware.NewLoggerMiddleware(router.Router)

	errorObj := http.ListenAndServe(addr, wrappedMux)

	log.Println("Started the server on the port: " + port)

	log.Fatal(errorObj)

}
