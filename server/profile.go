package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

// Profile : store user details
type Profile struct {
	ID, DisplayName, ImageURL, Email string
	torrents                         map[string]*Torrent
}

// return the user details
// null if not logged in
func userHandler(w http.ResponseWriter, r *http.Request) *appError {
	session, _ := SessionStore.Get(r, "default")

	js, err := json.Marshal(session.Values[googleProfileSessionKey])
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("content-type", "application/json")
	// w.Write(js)
	fmt.Fprintf(w, "%s", session.Values("Email"))
	return nil
}

// fetchProfile retrieves the Google+ profile of the user associated with the
// provided OAuth token.
func fetchProfile(ctx context.Context, tok *oauth2.Token) (*drive.About, error) {
	client := oauth2.NewClient(ctx, OAuthConfig.TokenSource(ctx, tok))
	// plusService, err := plus.New(client)
	driveService, err := drive.New(client)
	if err != nil {
		return nil, err
	}
	return driveService.About.Get().Fields("user/permissionId, user/photoLink, user/displayName, user/emailAddress").Do()
}

// stripProfile returns a subset of a drive.User.
func stripProfile(p *drive.About) *Profile {
	return &Profile{
		ID:          p.User.PermissionId,
		DisplayName: p.User.DisplayName,
		ImageURL:    p.User.PhotoLink,
		Email:       p.User.EmailAddress,
	}
}

// ProfileFromSession retreives the Gdrive profile from the default session.
// Returns nil if the profile cannot be retreived (e.g. user is logged out).
func ProfileFromSession(r *http.Request) *Profile {
	session, err := SessionStore.Get(r, defaultSessionID)
	if err != nil {
		return nil
	}
	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		return nil
	}
	profile, ok := session.Values[googleProfileSessionKey].(*Profile)
	if !ok {
		return nil
	}
	return profile
}
