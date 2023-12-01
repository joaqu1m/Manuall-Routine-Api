package route

import (
	"fmt"
	"manuall/routine-api/controller"
	"net/http"
	"strconv"
	"time"
)

var (
	crmOn      bool
	crmTicker  *time.Ticker
	crmTimeout = 15 * time.Second
)

func Crm() {
	http.HandleFunc("/crm/on", startRoutineCrm)
	http.HandleFunc("/crm/off", stopRoutineCrm)
	http.HandleFunc("/crm/check", checkRoutineCrm)
	http.HandleFunc("/crm/timeout", setTimeoutCrm)
}

func startRoutineCrm(w http.ResponseWriter, r *http.Request) {
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
	crmOn = false
	if w != nil {
		fmt.Fprintln(w)
	}
}

func checkRoutineCrm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, crmOn)
}

func setTimeoutCrm(w http.ResponseWriter, r *http.Request) {
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
