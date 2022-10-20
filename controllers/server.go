package controllers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/snoopz66/know-api-example/repositories"
	"github.com/urfave/negroni"
)

type Server struct {
	RootRouter *mux.Router
	Container  *repositories.Container
}

func Run() {
	s := &Server{}
	cont := &repositories.Container{}
	s.Container = cont
	s.RootRouter = initRouter(s)
	n := negroni.Classic()
	n.Use(getCORS())
	n.UseHandler(s.RootRouter)
	// Start server
	n.Run(":8080", "")
}

func initRouter(s *Server) *mux.Router {
	userController := User{UserService: s.Container.GetUserService()}
	RootRouter := mux.NewRouter().StrictSlash(false)
	RootRouter.HandleFunc("/users", userController.CreateUser).Methods("POST")
	RootRouter.HandleFunc("/users/{uuid}", userController.GetUser).Methods("GET")
	return RootRouter
}

func getCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Identity"},
		AllowCredentials: true,
	})
}
