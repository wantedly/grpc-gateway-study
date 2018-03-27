package app

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // db driver
	"github.com/wantedly/grpc-gateway-study/app/server"
)

// Run starts the grapiserver.
func Run() error {
	db, err := sqlx.Open("postgres", "postgres://postgres:@localhost/grpc-gateway-study?sslmode=disable")
	if err != nil {
		return err
	}
	defer db.Close()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(db),
		),
	)
	return s.Serve()
}
