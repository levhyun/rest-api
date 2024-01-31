package main

import (
	"flag"
	"rest-api/application/initializer"
	"rest-api/domain/user/domain/repository"
	"rest-api/domain/user/presentation"
	"rest-api/domain/user/presentation/router"
	"rest-api/domain/user/service"
	"strconv"
)

var port string

func init() {
	flogPort := flag.Int("p", 8080, "Enter the port")
	flag.Parse()
	port = ":" + strconv.Itoa(*flogPort)
}

func main() {
	db, err := initializer.NewDatabase()
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := presentation.NewUserController(userService)
	router.NewRouter(userController).Listen(port)
}
