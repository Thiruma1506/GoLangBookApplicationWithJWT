package main

import (
	config "bookapplication/Config"
	handlers "bookapplication/Handlers"
	repository "bookapplication/Repository"
	service "bookapplication/Service"
	"fmt"
	"log"
	"net/http"
)

func main() {

	//connect to the DB
	Database := config.ConnectToDB()

	//inject the DB to the repository
	repo := repository.NewMongoRepository(Database)

	//Inject the repo to the Handler
	bookHandler := &handlers.BookHandler{
		Repo: repo,
	}

	//=============================================

	userRepo := repository.NewUserRepository(Database)

	AuthService := service.NewAuthService(userRepo)

	userHandler := &handlers.UserHandler{
		AuthService: AuthService,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/books", bookHandler.CreateNewBook)
	mux.HandleFunc("PUT /api/books", bookHandler.UpdateBook)
	mux.HandleFunc("GET /api/books", bookHandler.GetBookById)
	mux.HandleFunc("DELETE /api/books", bookHandler.DeleteBookById)
	mux.HandleFunc("POST /api/list/books", bookHandler.CreateListOfBooks)

	mux.HandleFunc("POST /api/user/register",userHandler.RegisterNewUser)
	mux.HandleFunc("POST /api/user/login",userHandler.LoginUser)

	fmt.Print("Trying to start server on port 8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("The server start has failed %v", err)
	} else {
		fmt.Print("The server has successfully started on port 8080...")
	}
}
