package main

import (
	"URLshortener/controller"
	"URLshortener/repository"
	"URLshortener/service"
	"flag"
	"fmt"
)

func main() {
	urlRepository := getRepository()
	generator := service.NewCodeGenerator()
	urlService := service.NewUrlService(urlRepository, generator)
	handlers := controller.NewHandlers(urlService)
	server := controller.NewServer(handlers)
	server.Run()
}

func getRepository() repository.Repository {
	if getRepositoryFlag() {
		db := repository.NewDBRepository()
		err := db.Connect()
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Не удалось подключиться к бд, храним информацию в памяти")
			return repository.NewInMemoryRepository()
		} else {
			fmt.Println("Храним информацию в postgres")
			return db
		}
	} else {
		fmt.Println("Храним информацию в памяти")
		return repository.NewInMemoryRepository()
	}
}

func getRepositoryFlag() bool {
	repositoryFlag := flag.Bool("d", false, "Use db repository")
	flag.Parse()
	return *repositoryFlag
}
