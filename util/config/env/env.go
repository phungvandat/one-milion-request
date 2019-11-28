package env

import (
	"os"
)

// GetENV func
func GetENV() string {
	return os.Getenv("ENV")
}

// GetBackendPort func
func GetBackendPort() string {
	return os.Getenv("PORT")
}

// GetMaxQueue func
func GetMaxQueue() string {
	return os.Getenv("MAX_QUEUE")
}

// GetMaxWorker func
func GetMaxWorker() string {
	return os.Getenv("MAX_WORKER")
}

// GetMongoURI func
func GetMongoURI() string {
	return os.Getenv("MONGO_URI")
}

// GetMogoDBName func
func GetMogoDBName() string {
	return os.Getenv("MONGO_DB_NAME")
}
