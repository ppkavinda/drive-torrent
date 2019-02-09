package server

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"

	"github.com/ppkavinda/drive-torrent/profile"
)

const (
	oauthFlowRedirectKey    = "redirect"
	defaultSessionID        = "default"
	googleProfileSessionKey = "google_profile"
	oauthTokenSessionKey    = "oauth_token"
)

func init() {
	// Gob encoding for gorilla/sessions
	gob.Register(&oauth2.Token{})
	gob.RegisterName("*server.Profile", &profile.Profile{})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	sessionID := uuid.Must(uuid.NewV4()).String()
	oauthFlowSession, err := SessionStore.New(r, sessionID)
	if err != nil {
		fmt.Printf("could not create oauth session: %v\n", err)
	}
	oauthFlowSession.Options.MaxAge = 10 * 60 // 10 minutes

	redirectURL, err := validateRedirectURL(r.FormValue("redirect"))
	if err != nil {
		fmt.Printf("invalid redirect URL: %v\n", err)
	}
	oauthFlowSession.Values[oauthFlowRedirectKey] = redirectURL

	if err := oauthFlowSession.Save(r, w); err != nil {
		fmt.Printf("could not save session: %v", err)
	}

	// // Use the session ID for the "state" parameter.
	// // This protects against CSRF (cross-site request forgery).
	// // See https://godoc.org/golang.org/x/oauth2#Config.AuthCodeURL for more detail.
	url := OAuthConfig.AuthCodeURL(sessionID, oauth2.ApprovalForce,
		oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
	// fmt.Println("redirectURL", redirectURL)
}

// validateRedirectURL checks that the URL provided is valid.
// If the URL is missing, redirect the user to the application's root.
// The URL must not be absolute (i.e., the URL must refer to a path within this
// application).
func validateRedirectURL(path string) (string, error) {
	if path == "" {
		return "/", nil
	}

	// Ensure redirect URL is valid and not pointing to a different server.
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "/", err
	}
	if parsedURL.IsAbs() {
		return "/", errors.New("URL must not be absolute")
	}
	return path, nil
}

// logoutHandler clears the default session.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := SessionStore.New(r, defaultSessionID)
	if err != nil {
		fmt.Printf("could not get default session: %v\n", err)
	}
	session.Options.MaxAge = -1 // Clear session.
	profile.User = nil
	if err := session.Save(r, w); err != nil {
		fmt.Printf("could not save session: %v\n", err)
	}
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/"
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// oauthCallbackHandler completes the OAuth flow, retreives the user's profile
// information and stores it in a session.
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	oauthFlowSession, err := SessionStore.Get(r, r.FormValue("state"))
	if err != nil {
		fmt.Printf("invalid state parameter. try logging in again. %+v\n", err)
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)
	// Validate this callback request came from the app.
	if !ok {
		fmt.Printf("invalid state parameter. try logging in again. %+v\n", err)
	}

	code := r.FormValue("code")
	tok, err := OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("could not get auth token: %v\n", err)
	}

	session, err := SessionStore.Get(r, defaultSessionID)
	if err != nil {
		fmt.Printf("could not get default session: %v\n", err)
	}

	ctx := context.Background()
	profile, err := fetchProfile(ctx, tok)
	if err != nil {
		fmt.Printf("could not fetch Google profile: %v\n", err)
	}

	session.Values[oauthTokenSessionKey] = tok
	// Strip the profile to only the fields we need. Otherwise the struct is too big.
	session.Values[googleProfileSessionKey] = stripProfile(profile)
	if err := session.Save(r, w); err != nil {
		fmt.Printf("could not save session: %v\n", err)
	}

	tokenFile := profile.User.EmailAddress + ".json"
	saveToken(path.Join("./tokens", tokenFile), tok)

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Token: %+v\n", token)
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
