package server

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ppkavinda/drive-torrent/profile"
)

// here lies all the routes of the app
func registerRoutes(s *Server) {
	r := mux.NewRouter()

	// r.Handle("/", http.RedirectHandler("/ts", http.StatusFound))
	r.Path("/").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			w.Write([]byte("Hi"))
			return nil
		},
	))
	r.Methods("GET").Path("/books").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			fmt.Fprint(w, ProfileFromSession(r))
			return nil
		},
	))
	// r.Methods("GET").Path("/drive").Handler(appHandler(
	// 	func(w http.ResponseWriter, r *http.Request) *appError {
	// 		driveSample()
	// 		return nil
	// 	},
	// ))
	r.Methods("GET").Path("/ts").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			fmt.Printf("INFO %v\n", s.engine.GetTorrents())
			return nil
		},
	))
	r.Methods("GET").Path("/new").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			tmpl, err := template.ParseFiles("./index.html")
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return nil
			}
			tmpl.Execute(w, profile.User)
			return nil
		},
	))
	r.Methods("POST").Path("/new/magnet").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			magnet := r.FormValue("magnet")

			s.engine.NewMagnet(magnet, GetUser().Email)
			// fmt.Fprintf(w, "INFO4 %v", magnet)
			return nil
		},
	))
	r.Methods("POST").Path("/new/url").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			url := r.FormValue("url")
			remote, err := http.Get(url)
			if err != nil {
				fmt.Printf("Invalud remote url : %+v\n", err)
				return nil
			}
			fileData, err := ioutil.ReadAll(remote.Body)
			if err != nil {
				fmt.Printf("Failed To download remote torrent: %+v\n", err)
				return nil
			}

			reader := bytes.NewBuffer(fileData)
			info, err := metainfo.Load(reader)
			if err != nil {
				fmt.Printf("Unable to read metaInfo : %+v\n", err)
				return nil
			}

			spec := torrent.TorrentSpecFromMetaInfo(info)
			if err := s.engine.NewTorrentFromSpec(spec); err != nil {
				fmt.Printf("Torrent Error: %+v\n", err)
				return nil
			}
			return nil
		},
	))
	r.Methods("GET").Path("/view/{hash}").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			params := mux.Vars(r)

			s.engine.GetFiles(params["hash"])
			// fmt.Fprintf(w, "INFO4 %v", magnet)
			return nil
		},
	))
	// The following handlers are defined in auth.go and used in the
	// "Authenticating Users" part of the Getting Started guide.
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

	// [START request_logging]
	// Delegate all of the HTTP routing and serving to the gorilla/mux router.
	// Log all requests using the standard Apache format.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	// [END request_logging]

	r.Use(isLoggedIn)
}

func isLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Request = r
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
