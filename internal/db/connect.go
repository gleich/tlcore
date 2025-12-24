package db

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"go.mattglei.ch/timber"
)

func Connect() (*pgx.Conn, error) {
	const urlEnvName = "DATABASE_URL"
	url := os.Getenv(urlEnvName)
	if strings.TrimSpace(url) == "" {
		timber.FatalMsgf("failed to get postgres database url (%s)", urlEnvName)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("%w failed to connect to database", err)
	}
	return conn, nil
}
