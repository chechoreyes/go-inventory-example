package main

import (
	"context"

	"github.com/chechoreyes/max-inventory/database"
	"github.com/chechoreyes/max-inventory/settings"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {

	//Dependences injection
	app := fx.New(
		//Todas las funciones que devuelvan un Struct
		fx.Provide(context.Background(), settings.New, database.New),
		//Pasar comandos justo antes que empiece a correr
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USERS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
