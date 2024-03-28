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

func ValidateToken(fb *firebase.App, authToken string) (*auth.Token, error) {
	client, err := fb.Auth(context.Background())
	if err != nil {
		logger.Error("client error: ", err)
		return nil, err
	}
	token, err := client.VerifyIDToken(context.Background(), authToken)
	if err != nil {
		logger.Error("validation error: ", err)
		return nil, err
	}
	return token, nil
}
