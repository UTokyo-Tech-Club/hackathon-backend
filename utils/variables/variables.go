package variables

import (
	"hackathon-backend/utils/logger"
	"os"
)

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		logger.Fatal("Environment variable not set: ", k)
	}
	return v
}
