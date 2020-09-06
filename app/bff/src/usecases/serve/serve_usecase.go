package serve

import (
	"fmt"
	"io"
	"net/http"
)

type ServeUsecase struct {
	CurrentStatsHandler HandlerInterface
	SummaryStatsHandler HandlerInterface
}

func (httpServer *ServeUsecase) Serve(listen string) error {
	defaultHandler := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("hello")
		io.WriteString(w, "hello\n")
	}
	http.HandleFunc("/", defaultHandler)
	http.Handle("/current", httpServer.CurrentStatsHandler)
	http.Handle("/summary", httpServer.SummaryStatsHandler)

	fmt.Printf("listening: [%s]\n", listen)

	if err := http.ListenAndServe(listen, nil); err != nil {
		return err
	}
	return nil
}
