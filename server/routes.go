package server

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// here lies all the routes of the app
func registerRoutes(s *Server, r *mux.Router) *mux.Router {

	// r.Handle("/", http.RedirectHandler("/ts", http.StatusFound))
	r.Path("/").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			w.Write([]byte("Hi"))
			return nil
		},
	))

	r.Methods("GET").Path("/get").Handler(appHandler(s.getTorrentsHandler))
	r.Methods("GET").Path("/new").Handler(appHandler(newTorrentFormHandler))
	r.Methods("POST").Path("/new/magnet").Handler(appHandler(s.newMagnetHandler))
	r.Methods("POST").Path("/new/url").Handler(appHandler(s.newURLHandler))

	// r.Methods("GET").Path("/view/{hash}").Handler(appHandler(
	// 	func(w http.ResponseWriter, r *http.Request) *appError {
	// 		params := mux.Vars(r)

	// 		s.engine.GetFiles(params["hash"])
	// 		// fmt.Fprintf(w, "INFO4 %v", magnet)
	// 		return nil
	// 	},
	// ))

	r.Methods("GET").Path("/user").Handler(appHandler(userHandler))

	r.Methods("GET").Path("/login").Handler(appHandler(loginHandler))
	r.Methods("GET").Path("/logout").Handler(appHandler(logoutHandler))
	r.Methods("GET").Path("/oauth2callback").Handler(appHandler(oauthCallbackHandler))

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

	r.Use(isLoggedIn)
	return r
}

// Middleware
func isLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Request = r
		// enableCors(&w)
		if r.RequestURI != "/" && r.RequestURI != "/login" && !strings.HasPrefix(r.RequestURI, "/oauth2callback") {
			if ProfileFromSession(r) != nil {
				next.ServeHTTP(w, r)
			} else {
				http.Redirect(w, r, "/login", http.StatusUnauthorized)
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
