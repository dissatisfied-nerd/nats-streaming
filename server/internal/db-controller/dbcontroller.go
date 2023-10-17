package dbcontroller

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/jackc/pgx/v5"
)

type DBClient struct {
	conn sql.DB
	err  error
}

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")

	Client *DBClient
	once   sync.Once
)

func createClient() {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	conn, err := pgx.Connect(context.Background(), dbUrl)
	checkerror.CheckErr(err)

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	checkerror.CheckErr(err)
}

func NewDBClient() (*DBClient, error) {

	return Client, nil
}
