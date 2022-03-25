package router

import (
	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

	return router
}

// výstup z funkce se bude exportovat -> return router
// *mux.Router -> nadefinovaná reference na strukturu Routeru vytažená z mux balíčku
// http.HandleFunc funguje tak, že při spuštění zadané cesty v první části se správnou metodou definovanou
// v methods odpálí připojednou funkci naimportovatnou ze složky controller
