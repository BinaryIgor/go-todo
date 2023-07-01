package test

import (
	"fmt"
	"go-todo/app"
	"go-todo/shared"
)

func StartApp(modules ...shared.AppModule) int {
	port := RandomPort()
	app.Start(app.AppOptions{
		Address: fmt.Sprintf(":%d", port),
	}, modules...)
	return port
}
