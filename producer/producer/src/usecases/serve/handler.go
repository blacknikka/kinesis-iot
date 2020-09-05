package serve

import (
	"bytes"
	"fmt"
	"net/http"
)

type Serve struct {
	PostToIoT func(topic string, message string) error
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle")
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	err := s.PostToIoT("iot/stats", body)
	if err != nil {
		fmt.Printf("Error occurred : %v\n", err.Error())
	}
}
