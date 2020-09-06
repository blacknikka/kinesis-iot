package serve

import "net/http"

// HandlerInterface is a handler
type HandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
