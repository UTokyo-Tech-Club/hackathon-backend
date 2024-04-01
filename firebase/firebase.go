package firebaseAuth

import (
	"context"
	"hackathon-backend/utils/logger"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func Init() *firebase.App {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT")))
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error("Error initializing firebase: ", err)
	}
	logger.Info("Firebase initialized")
	return fb
}

func ValidateToken(fb *firebase.App, token string) (*auth.Token, error) {

	client, err := fb.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	idToken, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, err
	}

	return idToken, nil
}
