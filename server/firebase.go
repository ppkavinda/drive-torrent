package server

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// SaveDriveLink will save the (movie)drive link at firebase db
func SaveDriveLink(movieID, link, name string) {
	ctx := context.Background()

	opt := option.WithCredentialsFile("ytsag-e857b-firebase-adminsdk-shrm5-50421f9d81.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Printf("error initializing client: %v", err)
	}
	defer client.Close()

	hashes := client.Collection("hashes")
	hash := hashes.Doc(movieID)
	_, err = hash.Set(ctx, map[string]interface{}{
		name: map[string]interface{}{
			"link": link,
		},
	}, firestore.MergeAll)
	if err != nil {
		fmt.Printf("firebase get Error: %+v\n", err)
	}
}
