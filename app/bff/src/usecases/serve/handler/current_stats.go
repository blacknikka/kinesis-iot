package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/blacknikka/kinesis-iot/usecases/stats/current"
)

type ServeCurrentStats struct {
	CurrentUsecase current.CurrentStatsInterface
}

func (s ServeCurrentStats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count, err := s.CurrentUsecase.GetCurrentStartAmount("ver1")
	if err != nil {
		fmt.Printf("Error occurred : %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := current.CurrentStatsResponse{
		Kind:  "current",
		Stats: count,
	}

	content, err := json.Marshal(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(content))
}
