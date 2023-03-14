package DB

import (
	"blog-server-app/DB/entities"
	"fmt"
	"log"

	config "github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() *gorm.DB {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.GetString("database.host"),
		config.GetString("database.username"),
		config.GetString("database.password"),
		config.GetString("database.name"),
		config.GetInt32("database.port"))

	log.Println("Trying to establish db connection", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to the database")
	}

	//Migrate the models
	db.AutoMigrate(&entities.Blog{}, &entities.User{}, &entities.Comment{})

	log.Println("Connected to database successfully")

	return db
}
