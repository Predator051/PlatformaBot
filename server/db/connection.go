package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

var Ctx = context.Background()

func NewConn() (*pgx.Conn, error) {
	return pgx.Connect(Ctx, "postgres://tjiuxyka:DJHsolxYxZ4siRHxnAwFzkDYXYuaSBGb@dumbo.db.elephantsql.com/tjiuxyka")
}
