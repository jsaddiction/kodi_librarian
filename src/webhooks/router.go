package webhooks

import (
	radarr "kodi_librarian/src/webhooks/radarr"
	"net/http"

	"github.com/gorilla/mux"
)

// Ingres router
func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// Apply middleware
	router.Use(basicAuthMiddleware)

	// Define routes for each service
	router.HandleFunc("/radarr", radarr.HandleEvent).Methods(http.MethodPost)

	return router
}

// Middleware
func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate basic auth
		u, p, ok := r.BasicAuth()
		if !ok || u != "admin" || p != "admin" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
