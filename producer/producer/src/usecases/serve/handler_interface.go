package serve

import "net/http"

// IoTHandlerInterface is a handler in order to push to IoT Core
type IoTHandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// IoTResponse is a response for IoTHandlerInterface.
// This response is used for JSON response.
type IoTResponse struct {
	Message     string `json:"message"`
	WriteString string `json:"write_string"`
}
