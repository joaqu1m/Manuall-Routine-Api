package route

import (
	"fmt"
	"manuall/routine-api/controller"
	"net/http"
	"strconv"
	"time"
)

var (
	pipefyOn      bool
	pipefyTicker  *time.Ticker
	pipefyTimeout = 15 * time.Second
)

func Pipefy() {
	http.HandleFunc("/pipefy/on", startRoutinePipefy)
	http.HandleFunc("/pipefy/off", stopRoutinePipefy)
	http.HandleFunc("/pipefy/check", checkRoutinePipefy)
	http.HandleFunc("/pipefy/timeout", setTimeoutPipefy)
}

func startRoutinePipefy(w http.ResponseWriter, r *http.Request) {
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
	pipefyOn = false
	if w != nil {
		fmt.Fprintln(w)
	}
}

func checkRoutinePipefy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, pipefyOn)
}

func setTimeoutPipefy(w http.ResponseWriter, r *http.Request) {
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
