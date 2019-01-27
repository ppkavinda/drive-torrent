package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// here lies all the routes of the app
func registerRoutes(s *Server) {
	r := mux.NewRouter()

	r.Handle("/", http.RedirectHandler("/books", http.StatusFound))
	r.Methods("GET").Path("/books").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			fmt.Fprint(w, profileFromSession(r))
			return nil
		},
	))
	r.Methods("GET").Path("/drive").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			driveSample()
			return nil
		},
	))
	r.Methods("GET").Path("/ts").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			fmt.Printf("INFO %v\n", s.engine.GetTorrents())
			return nil
		},
	))
	r.Methods("POST").Path("/get").Handler(appHandler(
		func(w http.ResponseWriter, r *http.Request) *appError {
			magnet := r.FormValue("magnet") 
			// s.engine.NewMagnet("magnet:?xt=urn:btih:187D2FA2CD25E7256BC2101B8CAC2EDAEC82994B&dn=Quick+Heal+Total+Security+key+%28Renewal+for+1+year%29+%5BRxV%5D&tr=http%3A%2F%2Ffr33dom.h33t.com%3A3310%2Fannounce&tr=http%3A%2F%2Fwww.cbtorrents.com%3A2710%2Fannounce&tr=http%3A%2F%2Fnemesis.1337x.org%2Fannounce&tr=http%3A%2F%2Fexodus.1337x.org%2Fannounce&tr=http%3A%2F%2Ftracker.thepiratebay.org%2Fannounce&tr=http%3A%2F%2Ftracker.thepiratebay.org%3A80%2Fannounce&tr=http%3A%2F%2Fgenesis.1337x.org%3A1337%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.1337x.org%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce")
			// s.engine.NewMagnet("magnet:?xt=urn:btih:5FF320ED3B9DD2A06FA088E1DC5AE902FA8A0BC7&dn=All+Version+Office+Keys+by+mansory.txt&tr=http%3A%2F%2Fbt.rghost.net%2Fannounce&tr=http%3A%2F%2Ftracker.publicbt.com%2Fannounce&tr=udp%3A%2F%2Ftracker.publicbt.com%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.openbittorrent.com%2Fannounce&tr=udp%3A%2F%2Ftracker.1337x.org%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce")
			s.engine.NewMagnet(magnet)
			// fmt.Fprintf(w, "INFO4 %v", magnet)
			return nil
		},
	))
	// The following handlers are defined in auth.go and used in the
	// "Authenticating Users" part of the Getting Started guide.
	r.Methods("GET").Path("/login").Handler(appHandler(loginHandler))
	r.Methods("POST").Path("/logout").Handler(appHandler(logoutHandler))
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
}
