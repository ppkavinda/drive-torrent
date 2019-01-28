package profile

// Profile : store user details
type Profile struct {
	ID, DisplayName, ImageURL, Email string
	// torrents                         map[string]*Torrent
}

// User : stores user object globally
var User *Profile
