package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// here lies all the routes of the app
func getRoutes(s *Server, r *mux.Router) *mux.Router {

	r.Handle("/", http.FileServer(http.Dir("./static")))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// http.Handle("/new", IsLoggedIn(r))

	r.Methods("GET").Path("/get").HandlerFunc(s.getTorrentsHandler)
	r.Methods("GET").Path("/new").HandlerFunc(newTorrentFormHandler)
	r.Methods("POST").Path("/new/magnet").HandlerFunc(s.newMagnetHandler)
	r.Methods("POST").Path("/new/url").HandlerFunc(s.newURLHandler)

	r.Methods("GET").Path("/user").HandlerFunc(userHandler)

	r.Methods("GET").Path("/login").HandlerFunc(loginHandler)
	r.Methods("GET").Path("/logout").HandlerFunc(logoutHandler)
	r.Methods("GET").Path("/oauth2callback").HandlerFunc(oauthCallbackHandler)

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	// add middleware
	http.Handle("/get", IsLoggedIn(r))
	http.Handle("/new", IsLoggedIn(r))
	http.Handle("/new/url", IsLoggedIn(r))
	http.Handle("/new/magnet", IsLoggedIn(r))
	http.Handle("/user", IsLoggedIn(r))

	return r
}

// IsLoggedIn :  Middleware
func IsLoggedIn(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ProfileFromSession(r) == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
