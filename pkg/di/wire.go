//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "githum.com/athunlal/bookNowAdmin-svc/pkg/api"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/api/handler"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/config"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/db"
	repoesitory "githum.com/athunlal/bookNowAdmin-svc/pkg/repository"
	usecas "githum.com/athunlal/bookNowAdmin-svc/pkg/usecase"
)

func InitApi(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repoesitory.NewAdminRepo,
		usecas.NewAdminUseCase,
		usecas.NewJWTuseCase,
		handler.NewAdminHandler,
		http.NewServerHttp,
	)

	return &http.ServerHttp{}, nil
}
