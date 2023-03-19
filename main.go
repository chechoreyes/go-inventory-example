package main

import (
	"context"

	"github.com/chechoreyes/max-inventory/database"
	"github.com/chechoreyes/max-inventory/internal/repository"
	"github.com/chechoreyes/max-inventory/internal/service"
	"github.com/chechoreyes/max-inventory/settings"
	"go.uber.org/fx"
)

func main() {

	//Dependences injection
	app := fx.New(
		//Todas las funciones que devuelvan un Struct
		fx.Provide(context.Background, settings.New, database.New, repository.New, service.New),
		//Pasar comandos justo antes que empiece a correr
		fx.Invoke(
			// func(ctx context.Context, serv service.Service) {
			// 	err := serv.RegisterUser(ctx, "my@email.com", "myname", "mypassword")
			// 	if err != nil {
			// 		panic(err)
			// 	}

			// 	u, err := serv.LoginUser(ctx, "my@email.com", "mypassword")
			// 	if err != nil {
			// 		panic(err)
			// 	}

			// 	if u.Name != "myname" {
			// 		panic("wrong name")
			// 	}
			// },
		),
	)

	app.Run()
}
