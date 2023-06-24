package main

import (
	"context"
	"errors"
	"fmt"
	"go-todo/shared"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

type CurrentUser struct {
	id uuid.UUID
}

type AppOptions struct {
	address string
}

const USER_KEY = "user"

type Middleware func(http.Handler) http.Handler

func recoveryMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in middleware", r)
					shared.WriteJsonResponse(w, 500, struct {
						Name string
					}{"Error"})
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}

func loggingMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Before request")
			defer log.Println("After request....")
			h.ServeHTTP(w, r)
		})
	}
}

func authMiddleware() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Values("Authorization")
			if len(authHeader) == 0 {
				authHeader = r.URL.Query()["Authorization"]
			}
			//TODO: real checks!
			if len(authHeader) == 0 {
				w.WriteHeader(401)
			} else {
				h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), USER_KEY,
					CurrentUser{uuid.New()})))
			}
		})
	}
}

func newHandler(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

var server *http.Server

func main() {
	fmt.Println("Starting TODO app!")

	Start(AppOptions{":8080"})

	fmt.Println("Waiting for ending signal...")

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	<-signChan

	fmt.Println("Signal received, stopping server...")

	Stop()
}

func Start(options AppOptions) {
	server = &http.Server{Addr: options.address}

	registerHandler("/", func(w http.ResponseWriter, r *http.Request) {
		currentUser := r.Context().Value(USER_KEY)
		fmt.Println("Current user: ", currentUser)
		// panic("test")
		shared.WriteJsonOkResponse(w, struct {
			Name string
		}{"User"})

	})

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Http server error: %v", err)
		}
		log.Println("Stopped serving new connections")
	}()
}

func registerHandler(path string, h http.HandlerFunc) {
	http.Handle(path, newHandler(h, recoveryMiddleware(), loggingMiddleware(),
		authMiddleware()))
}

func Stop() {
	if server == nil {
		return
	}

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete")
}
