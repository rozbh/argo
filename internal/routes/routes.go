package routes

import (
	"net/http"

	"github.com/rozbh/argo/internal/controllers"
)

// InitRoutes sets up the application's URL routes.
func InitRoutes() http.Handler {
	mux := http.NewServeMux()

	// Homepage route
	mux.HandleFunc("/", controllers.HomeHandler)

	return mux
}
