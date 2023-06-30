package test

import (
	"go-todo/app"
	"go-todo/shared"
)

func StartApp(modules ...shared.AppModule) int {
	port := RandomPort()
	app.Start(app.AppOptions{
		Address: ":" + string(port),
	}, modules...)
	return port
}
