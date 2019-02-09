package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// here lies all the routes of the app
func getRoutes(s *Server, r *mux.Router) *mux.Router {

	r.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("index.html").Funcs(template.FuncMap{
			"marshal": func(v interface{}) string {
				a, _ := json.Marshal(v)
				return string(a)
			},
		}).ParseFiles("static/index.html"))
		t.Execute(w, GetUser(r))
		// fmt.Printf("%+v\n", t.Execute(w, GetUser()))
	})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	r.Methods("GET").Path("/get").HandlerFunc(s.getTorrentsHandler)
	r.Methods("GET").Path("/new").HandlerFunc(newTorrentFormHandler)
	r.Methods("POST").Path("/new/magnet").HandlerFunc(s.newMagnetHandler)
	r.Methods("POST").Path("/new/url").HandlerFunc(s.newURLHandler)

	r.Methods("POST").Path("/torrent/stop").HandlerFunc(s.stopTorrent)
	r.Methods("POST").Path("/torrent/remove").HandlerFunc(s.removeTorrent)
	r.Methods("POST").Path("/torrent/start").HandlerFunc(s.startTorrent)

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

	http.Handle("/remove", IsLoggedIn(r))
	// http.Handle("/user", IsLoggedIn(r))

	r.Use(RegisterRequest)
	return r
}

// IsLoggedIn :  Middleware
func IsLoggedIn(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ProfileFromSession(r) == nil {
			if reqType := r.Header.Get("X-Requested-With"); reqType == "xmlhttprequest" {
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("{}"))
			} else {
				http.Redirect(w, r, "/login", http.StatusFound)
			}
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

// RegisterRequest : register request as a global
func RegisterRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Request = r
		next.ServeHTTP(w, r)
	})

}
