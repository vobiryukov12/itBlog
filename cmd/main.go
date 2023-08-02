package main

import (
	itblog "github.com/Zemavong/itBlog"
	"github.com/Zemavong/itBlog/pkg/handler"
	"github.com/Zemavong/itBlog/pkg/repository"
	"github.com/Zemavong/itBlog/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//fileServer := http.FileServer(http.Dir("./template"))

	//http.HandleFunc("/hi", handler.IndexHandler)

	db, err := repository.NewMysqlDB(repository.Config{
		User:   viper.GetString("username"),
		Passwd: viper.GetString("password"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "ItBlog",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(itblog.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	logrus.Print("TodoApp Started")

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
