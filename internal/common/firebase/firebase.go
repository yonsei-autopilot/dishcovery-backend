package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client

func InitializeFirebaseClient() {
	ctx := context.Background()

	sa := option.WithCredentialsFile("firebase-application-credentials.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		panic(fmt.Sprintf("Failed to create Firebase app: %v", err))
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Firestore: %v", err))
	}

	firestoreClient = client
}

func GetClient() *firestore.Client {
	return firestoreClient
}
