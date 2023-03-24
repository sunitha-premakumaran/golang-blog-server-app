package main

import (
	"blog-server-app/DB"
	middleware "blog-server-app/modules/system/middlewares"
	router "blog-server-app/routes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	appLogger := logger.Named("main")

	port := config.GetString("port")

	if port == "" {
		appLogger.Info("No PORT set. Setting it to default 3000")
		port = "3000"
	}

	addr := fmt.Sprintf(":%s", port)

	//Initiate db connection
	db := DB.New(logger.Named("main/db")).InitConnection()

	//Initialize routes
	router := router.NewRouter(db, logger.Named("main/router"))

	loggerMiddleware := middleware.NewLoggerMiddleware(router.Router, appLogger)

	serverErrors := make(chan error, 1)

	go func() {
		serverErrors <- http.ListenAndServe(addr, loggerMiddleware)
	}()
	appLogger.Info("Started the server on the port: " + port)

	// Set callback for the signal interrupt
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		appLogger.Error(err.Error())
		return

	case sig := <-shutdown:
		connection, _ := db.DB()
		appLogger.Info("Recieved interrupt signal. Shutting down server and closing db connection")
		appLogger.Error(sig.String())
		connection.Close()

	}
}
