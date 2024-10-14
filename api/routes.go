package api

import "github.com/gorilla/mux"

func CreateRoutes(userHandler *handlers.userHandler, carsHandler *handlers.CarsHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/list", userHandler.List).Methods("GET")
	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET")

	r.HandleFunc("/users/create", carsHandler.Create).Methods("POST")
	r.HandleFunc("/users/list", carsHandler.List).Methods("GET")
	r.HandleFunc("/users/find/{id:[0-9]+}", carsHandler.Find).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
