package router

import (
	"github.com/gorilla/mux"
	contollers "github.com/kaiwalyakoparkar/practical-devops/tree/main/Languages/Go/Projects/controllers"
)

func Router() *mux.Router {
	routers := mux.NewRouter()

	routers.Handle("/api/movies", contollers.g)

	return routers
}