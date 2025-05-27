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
	r.RestServe.GET("/api/users/get/:user_id", handl.GetUser(ctx))
	r.RestServe.DELETE("/api/users/delete/:user_id", handl.DeleteUser(ctx))
	r.RestServe.POST("/api/users/update/:user_id", handl.UpdateUser(ctx))
	r.RestServe.POST("/api/users/find", handl.FindWithFilter(ctx))
}
