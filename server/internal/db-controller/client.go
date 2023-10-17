package dbcontroller

import "database/sql"

type DBClient struct {
	conn *sql.DB
	err  error
}
