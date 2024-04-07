package main

import (
	"github.com/MrClean-code/testD/pkg/handler"
	"github.com/MrClean-code/testD/pkg/repository"
	"github.com/MrClean-code/testD/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes()

	port := viper.GetString("PORT")

	addr := ":" + port
	if port == "" {
		log.Fatal("Port not specified in the configuration")
	}

	log.Printf("Server is running on port %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs") // C:\Users\Dc\Desktop\david\wbschool_exam_L2\develop\dev11\configs <-- если через терминал
	return viper.ReadInConfig()
}
