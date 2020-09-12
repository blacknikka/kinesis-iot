package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	summaryUsecase "github.com/blacknikka/kinesis-iot/usecases/stats/summary/get"
)

type ServeSummaryStats struct {
	SummaryUsecase summaryUsecase.SummaryStatsInterface
}

func (s ServeSummaryStats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sumaryContent, err := s.SummaryUsecase.GetSummaryStartAmount("ver1")
	if err != nil {
		fmt.Printf("Error occurred : %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := summaryUsecase.SummaryStatsResponse{
		Kind:    "current",
		Summary: sumaryContent,
	}

	content, err := json.Marshal(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(content))
}
