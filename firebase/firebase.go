package firebaseAuth

import (
	"context"
	"fmt"
	"hackathon-backend/utils/logger"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func Init() *firebase.App {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT")))
	logger.Info(os.Getenv("FIREBASE_SERVICE_ACCOUNT"))
	fb, err := firebase.NewApp(context.Background(), nil, opt)
	if fb == nil {
		logger.Error("Firebase app not created")
		return nil
	}
	ctx := context.Background()
	if _, err := fb.Auth(ctx); err != nil {
		logger.Error("Failed to establish connection with Firebase Auth: ", err)
		return nil
	}
	if err != nil {
		logger.Error("Error initializing firebase: ", err)
		return nil
	}
	logger.Info("Firebase initialized")
	return fb
}

func ValidateToken(fb *firebase.App, token string) (*auth.Token, error) {
	if fb == nil {
		return nil, fmt.Errorf("firebase app is nil")
	}

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
