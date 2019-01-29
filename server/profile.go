package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"

	"github.com/ppkavinda/drive-torrent/profile"
)

// return the user details
// null if not logged in
func userHandler(w http.ResponseWriter, r *http.Request) *appError {
	// session, _ := SessionStore.Get(r, "default")

	// js, err := json.Marshal(session.Values[googleProfileSessionKey])

	js, err := json.Marshal(GetUser())
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("content-type", "application/json")
	w.Write(js)
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
func stripProfile(p *drive.About) *profile.Profile {
	profile.User = &profile.Profile{
		ID:          p.User.PermissionId,
		DisplayName: p.User.DisplayName,
		ImageURL:    p.User.PhotoLink,
		Email:       p.User.EmailAddress,
	}
	fmt.Printf("%v", profile.User)
	return profile.User
}

// ProfileFromSession retreives the Gdrive profile from the default session.
// Returns nil if the profile cannot be retreived (e.g. user is logged out).
func ProfileFromSession(r *http.Request) *profile.Profile {
	session, err := SessionStore.Get(r, defaultSessionID)
	if err != nil {
		return nil
	}
	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		return nil
	}
	profile, ok := session.Values[googleProfileSessionKey].(*profile.Profile)
	if !ok {
		return nil
	}

	return profile
}

// GetUser : getter for User variable
// if nil check session or token file
func GetUser() *profile.Profile {
	if profile.User == nil {
		session, _ := SessionStore.Get(Request, "default")

		user := session.Values[googleProfileSessionKey]

		profile.User = &profile.Profile{
			ID:          user.(*profile.Profile).ID,
			DisplayName: user.(*profile.Profile).DisplayName,
			ImageURL:    user.(*profile.Profile).ImageURL,
			Email:       user.(*profile.Profile).Email,
		}
	}

	return profile.User
}
