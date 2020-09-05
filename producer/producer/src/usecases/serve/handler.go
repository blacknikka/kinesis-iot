package serve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Serve struct {
	PostToIoT func(topic string, message string) error
}

func (s *Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()

	err := s.PostToIoT("iot/stats", body)
	if err != nil {
		fmt.Printf("Error occurred : %v\n", err.Error())
		resp := IoTResponse{
			Message: fmt.Sprintf("Error occurred : %v", err.Error()),
		}

		content, err := json.Marshal(&resp)
		if err != nil {
			io.WriteString(w, `{"message": "something wrong"}`)
		}
		io.WriteString(w, string(content))

		return
	}

	resp := IoTResponse{
		Message:     "success",
		WriteString: body,
	}
	content, err := json.Marshal(&resp)
	if err != nil {
		io.WriteString(w, `{"message": "Write was success, but something wrong after writing."}`)
	}
	io.WriteString(w, string(content))
}
