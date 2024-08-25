package app

import (
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
	pg, err := postgres.NewPostgresPool(postgres.PostgresConfig{
		ConnectionString: cfg.PG.URL,
		MaxConns:         20,
	})
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepositories(pg)
	service := service.NewService(service.ServicesDependencies{
		Repos:    repo,
		SignKey:  cfg.SignKey,
		TokenTTL: cfg.TokenTTL,
		Salt:     cfg.Salt,
	})

	r := chi.NewRouter()
	handler := v1.NewHandler(service)
	server := v1.NewServer(handler, r)
	server.Router()

}
