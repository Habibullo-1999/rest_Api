package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Habibullo-1999/rest_Api/cmd/app"
	"github.com/Habibullo-1999/rest_Api/pkg/note"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/dig"

)

func main() {
	host := "0.0.0.0"
	port := "9999"
	dsn := "postgres://app:pass@localhost:5432/db"

	if err := execute(host, port, dsn); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func execute(host string, port string, dsn string) (err error) {

	deps := []interface{}{

		app.NewServer,
		mux.NewRouter,
		func() (*pgxpool.Pool, error) {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			return pgxpool.Connect(ctx, dsn)
		},
		note.NewService,
		func(s *app.Server) *http.Server {
			return &http.Server{
				Addr:    net.JoinHostPort(host, port),
				Handler: s,
			}
		},
	}

	container := dig.New()
	for _, dep := range deps {
		err = container.Provide(dep)
		if err != nil {
			return err
		}
	}

	err= container.Invoke(func(server *app.Server){
		server.Init()
	})
	if err != nil {
		return err
	}

	return container.Invoke(func (s *http.Server) error {
		return s.ListenAndServe()
	})

}
