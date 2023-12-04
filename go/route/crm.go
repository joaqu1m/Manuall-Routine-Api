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
	crmOn      bool
	crmTicker  *time.Ticker
	crmTimeout = 15 * time.Second
)

func Crm(router *mux.Router) {
	router.HandleFunc("/crm/on", startRoutineCrm).Methods("GET", "OPTIONS")
	router.HandleFunc("/crm/off", stopRoutineCrm).Methods("GET", "OPTIONS")
	router.HandleFunc("/crm/check", checkRoutineCrm).Methods("GET", "OPTIONS")
	router.HandleFunc("/crm/timeout", setTimeoutCrm).Methods("POST", "OPTIONS")
}

func startRoutineCrm(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	crmOn = true
	go func() {
		crmTicker = time.NewTicker(crmTimeout)
		for {
			if !crmOn {
				break
			}
			select {
			case <-crmTicker.C:
				controller.Crm()
			}
		}
	}()
	if w != nil {
		fmt.Fprintln(w)
	}
}

func stopRoutineCrm(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	crmOn = false
	if w != nil {
		fmt.Fprintln(w)
	}
}

func checkRoutineCrm(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintln(w, crmOn)
}

func setTimeoutCrm(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	timeoutStr := r.URL.Query().Get("timeout")
	if timeoutStr != "" {
		newTimeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			http.Error(w, "Valor inválido", http.StatusBadRequest)
			return
		}

		crmTimeout = time.Duration(newTimeout) * time.Second

		stopRoutineCrm(nil, nil)
		startRoutineCrm(nil, nil)
	}
	fmt.Fprintln(w)
}
