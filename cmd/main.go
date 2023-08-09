package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/handler"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/service"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	//var brands []carsBrandRandomGenerator.Brand
	//brands, _ = repos.Brand.GetAll()
	//fmt.Println(brands[0].Name, err)
	//imagePath := "download.png"
	//imageBytes, err := readImageFile(imagePath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//repos.Brand.Create(carsBrandRandomGenerator.Brand{
	//	Id:         1,
	//	Name:       "Audi",
	//	ImageBrand: imageBytes,
	//})
	handlers := handler.NewHandler(services)
	srv := new(carsBrandRandomGenerator.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func readImageFile(filename string) ([]byte, error) {
	imageBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return imageBytes, nil
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
