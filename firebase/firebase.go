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
	// isDeployed := os.Getenv("IS_DEPLOYED")

	// var opt option.ClientOption
	// if isDeployed == "true" {
	// 	opt = option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON")))
	// } else {
	// }
	// opt := option.WithCredentialsFile(os.Getenv("FIREBASE_SERVICE_ACCOUNT_JSON"))
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
	logger.Info("Client: ", client)
	if err != nil {
		logger.Error("client error: ", err)
		return nil, err
	}
	token, err := client.VerifyIDToken(context.Background(), authToken)
	logger.Info("Token: ", token)
	if err != nil {
		logger.Error("validation error: ", err)
		return nil, err
	}
	return token, nil
}
