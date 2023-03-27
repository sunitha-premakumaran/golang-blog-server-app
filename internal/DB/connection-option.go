package DB

import (
	"blog-server-app/internal/DB/entities"
	"fmt"

	config "github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppDB struct {
	logger *zap.Logger
}

func (dbInstance *AppDB) InitConnection() *gorm.DB {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.GetString("database.host"),
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.name"),
		config.GetInt32("database.port"))

	dbInstance.logger.Info("Trying to establish db connection")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to the database")
	}

	//Migrate the models
	errorDb := db.AutoMigrate(&entities.Blog{}, &entities.User{}, &entities.Comment{})

	if errorDb != nil {
		dbInstance.logger.Panic(errorDb.Error())
		panic(errorDb.Error())
	}

	dbInstance.logger.Info("Connected to database successfully")

	return db
}

func New(logger *zap.Logger) *AppDB {
	return &AppDB{logger}
}
