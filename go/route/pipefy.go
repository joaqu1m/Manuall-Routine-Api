package route

import (
	"fmt"
	"manuall/routine-api/controller"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

var (
	pipefyOn      bool
	pipefyTicker  *time.Ticker
	pipefyTimeout = 15 * time.Second
)

func Pipefy(router *mux.Router) {
	router.HandleFunc("/pipefy/on", startRoutinePipefy)
	router.HandleFunc("/pipefy/off", stopRoutinePipefy)
	router.HandleFunc("/pipefy/check", checkRoutinePipefy)
	router.HandleFunc("/pipefy/timeout", setTimeoutPipefy)
}

func startRoutinePipefy(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	pipefyOn = true
	go func() {
		pipefyTicker = time.NewTicker(pipefyTimeout)
		for {
			if !pipefyOn {
				break
			}
			select {
			case <-pipefyTicker.C:
				controller.Pipefy()
			}
		}
	}()
	if w != nil {
		fmt.Fprintln(w)
	}
}

func stopRoutinePipefy(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	pipefyOn = false
	if w != nil {
		fmt.Fprintln(w)
	}
}

func checkRoutinePipefy(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, pipefyOn)
}

func setTimeoutPipefy(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	timeoutStr := r.URL.Query().Get("timeout")
	if timeoutStr != "" {
		newTimeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			http.Error(w, "Valor inválido", http.StatusBadRequest)
			return
		}

		pipefyTimeout = time.Duration(newTimeout) * time.Second

		stopRoutinePipefy(nil, nil)
		startRoutinePipefy(nil, nil)
	}
	fmt.Fprintln(w)
}
