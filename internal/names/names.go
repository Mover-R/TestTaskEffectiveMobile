package names

import (
	"TestTaskEffectiveMobile/internal/names/handler"
	"TestTaskEffectiveMobile/internal/names/repository"
	"TestTaskEffectiveMobile/internal/names/service"
	server "TestTaskEffectiveMobile/internal/transport/rest"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(ctx context.Context, r server.Router, pgDB *pgxpool.Pool) {
	repo := repository.NewRepository(pgDB)
	serv := service.NewService(repo)
	handl := handler.NewHandler(serv)

	r.RestServe.PUT("/api/users/create", handl.CreateUser(ctx))
}
