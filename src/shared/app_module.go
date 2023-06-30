package shared

import "net/http"

type AppModule struct {
	Router http.Handler
	Path   string
}
