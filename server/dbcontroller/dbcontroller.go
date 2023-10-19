package dbcontroller

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/dissatisfied-nerd/nats-streaming/pkg/checkerror"
	"github.com/jackc/pgx/v5"
)

type DBClient pgx.Conn

var (
	dbUrl = os.Getenv("POSTGRES_URL")

	Client *DBClient
	once   sync.Once
)

func NewDBClient() (*DBClient, error) {
	once.Do(func() {
		fmt.Println(dbUrl)
		conn, err := pgx.Connect(context.Background(), dbUrl)
		checkerror.CheckErr(err)

		err = conn.Ping(context.Background())
		checkerror.CheckErr(err)
	})

	return Client, nil
}
