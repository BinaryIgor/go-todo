package main

import (
	"fmt"
	"go-todo/app"
	"go-todo/shared"
	"go-todo/user"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting TODO app!")

	userModule := user.Module(user.ModuleConfigFromEnvVariables(&shared.UtcClock))

	app.Start(app.AppOptions{Address: ":8080"}, userModule)

	fmt.Println("Waiting for ending signal...")

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	<-signChan

	fmt.Println("Signal received, stopping server...")

	app.Stop()
}
