package db

import (
	"fmt"
	"os"
	"strings"

	"go.mattglei.ch/timber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	const urlEnvName = "DATABASE_URL"
	url := os.Getenv(urlEnvName)
	if strings.TrimSpace(url) == "" {
		timber.FatalMsgf("failed to get postgres database url (%s)", urlEnvName)
	}

	conn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%w failed to connect to database", err)
	}
	return conn, nil
}
