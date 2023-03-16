package main

import (
	"log"

	"github.com/chechoreyes/max-inventory/settings"
	"go.uber.org/fx"
)

func main() {

	//Dependences injection
	app := fx.New(
		//Todas las funciones que devuelvan un Struct
		fx.Provide(settings.New),
		//Pasar comandos justo antes que empiece a correr
		fx.Invoke(
			func(s *settings.Settings) {
				log.Println(s)
			},
		),
	)

	app.Run()
}
