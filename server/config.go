package server

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// SessionStore : store session globally
	SessionStore sessions.Store
	// OAuthConfig : store authConfigs globally
	OAuthConfig *oauth2.Config
	// Request : make http request global
	Request *http.Request
)

func init() {
	var cookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore

	OAuthConfig = configureOAuthClient("404364039745-0caba0fvhaja2cogru4jvl0gqq3anf50.apps.googleusercontent.com", "zRly0iH-ThMZrYRxER5PT_ue")
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:8080/oauth2callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile", "https://www.googleapis.com/auth/drive"},
		Endpoint:     google.Endpoint,
	}
}
