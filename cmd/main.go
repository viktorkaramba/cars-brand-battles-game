package main

import (
	"github.com/spf13/viper"
	carsBrandRandomGenerator "github.com/viktorkaramba/cars-brand-random-generator-app"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/handler"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/repository"
	"github.com/viktorkaramba/cars-brand-random-generator-app/pkg/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(carsBrandRandomGenerator.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
