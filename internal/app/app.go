package app

import (
	"kode-notes/internal/config"
	v1 "kode-notes/internal/controller/http/v1"
	"kode-notes/internal/repository"
	"kode-notes/internal/service"
	"kode-notes/internal/spellchecker"
	"kode-notes/pkg/postgres"
	"log/slog"
	"os"

	"github.com/go-chi/chi"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run(path string) {
	cfg, err := config.NewConfig(path)
	if err != nil {
		panic(err)
	}

	log := setupLogger(cfg.Level)
	log = log.With(slog.String("env", cfg.Level))
	log.Info("initializing server", slog.String("address", cfg.HTTP.Host+":"+cfg.HTTP.Port))
	log.Debug("logger debug mode enabled")

	pg, err := postgres.NewPostgresPool(postgres.PostgresConfig{
		ConnectionString: cfg.PG.URL,
		MaxConns:         20,
	})
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepositories(pg)
	yandexspeller := spellchecker.NewYandexSpellChecker()
	service := service.NewService(service.ServicesDependencies{
		Repos:    repo,
		SignKey:  cfg.SignKey,
		TokenTTL: cfg.TokenTTL,
		Salt:     cfg.Salt,
		Speller:  yandexspeller,
	})

	log.Info("starting server", slog.String("address", cfg.HTTP.Host+":"+cfg.HTTP.Port))
	r := chi.NewRouter()
	handler := v1.NewHandler(service)
	server := v1.NewServer(handler, r)
	server.Router()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
