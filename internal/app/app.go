package app

import (
	"context"
	"kode-notes/internal/config"
	v1 "kode-notes/internal/controller/http/v1"
	"kode-notes/internal/repository"
	"kode-notes/internal/service"
	"kode-notes/pkg/postgres"

	"github.com/go-chi/chi"
)

func Run(path string) {

	cfg, err := config.NewConfig(path)
	if err != nil {
		panic(err)
	}
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		panic(err)
	}
	err = pg.Pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepositories(pg)
	deps := service.ServicesDependencies{
		Repos:    repo,
		SignKey:  cfg.SignKey,
		TokenTTL: cfg.TokenTTL,
	}
	service := service.NewService(deps)
	r := chi.NewRouter()
	handler := v1.NewHandler(service)
	server := v1.NewServer(handler, r)
	server.Router()

}
