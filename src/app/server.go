package app

import (
	"context"

	"errors"
	"fmt"
	"go-todo/shared"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

type CurrentUser struct {
	id uuid.UUID
}

type AppOptions struct {
	Address string
}

const USER_KEY = "user"

type Middleware func(http.Handler) http.Handler

func recoveryMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in middleware", r)
				shared.WriteJsonErrorResponse(w, r)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

// func loggingMiddleware() Middleware {
// 	return func(h http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			log.Println("Before request")
// 			defer log.Println("After request....")
// 			h.ServeHTTP(w, r)
// 		})
// 	}
// }

func authMiddleware_() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Values("Authorization")
			if len(authHeader) == 0 {
				authHeader = r.URL.Query()["Authorization"]
			}
			publicEndpoint := shared.IsEndpointPublic(r.URL.Path)
			//TODO: real checks!
			if !publicEndpoint && len(authHeader) == 0 {
				w.WriteHeader(401)
			} else {
				h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), USER_KEY,
					CurrentUser{uuid.New()})))
			}
		})
	}
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Values("Authorization")
		if len(authHeader) == 0 {
			authHeader = r.URL.Query()["Authorization"]
		}
		publicEndpoint := shared.IsEndpointPublic(r.URL.Path)
		//TODO: real checks!
		if !publicEndpoint && len(authHeader) == 0 {
			w.WriteHeader(401)
		} else {
			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), USER_KEY,
				CurrentUser{uuid.New()})))
		}
	})
}

func newHandler(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

var server *http.Server

func Start(options AppOptions, modules ...shared.AppModule) {
	router := chi.NewRouter()

	server = &http.Server{Addr: options.Address, Handler: router}

	//TODO: proper cors config depending on the env
	corsHandler := cors.AllowAll().Handler

	router.Use(corsHandler, middleware.Logger, middleware.Timeout(60*time.Second),
		recoveryMiddleware, authMiddleware)

	for _, m := range modules {
		router.Mount(m.Path, m.Router)
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Http server error: %v", err)
		}
		log.Println("Stopped serving new connections")
	}()
}

// func registerHandler(path string, h http.HandlerFunc) {
// 	http.Handle(path, newHandler(h, recoveryMiddleware(), loggingMiddleware(),
// 		authMiddleware()))
// }

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
