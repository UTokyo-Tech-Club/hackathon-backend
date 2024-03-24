package firebaseAuth

import (
	"context"
	"hackathon-backend/utils/logger"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	logger := logger.NewLogger()

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON"))
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Fatal("Error initializing firebase: ", err)
	}
	logger.Info("Firebase initialized")
	return fb
}

func ValidateToken(fb *firebase.App, authToken string) (*auth.Token, error) {
	client, err := fb.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	token, err := client.VerifyIDToken(context.Background(), authToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
