package route

import (
	"fmt"
	"manuall/routine-api/consumable"
	"manuall/routine-api/controller"
	"net/http"
	"strconv"
	"time"
)

var (
	crmOn      bool
	crmTicker  *time.Ticker
	crmTimeout = 3 * time.Second
)

func Crm() {
	consumable.Crm("joaquim.pires@sptech.school")
	controller.Crm()
	http.HandleFunc("/crm/on", startRoutine)
	http.HandleFunc("/crm/off", stopRoutine)
	http.HandleFunc("/crm/check", checkRoutine)
	http.HandleFunc("/crm/timeout", setTimeout)
}

func startRoutine(w http.ResponseWriter, r *http.Request) {
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

func stopRoutine(w http.ResponseWriter, r *http.Request) {
	crmOn = false
	if w != nil {
		fmt.Fprintln(w)
	}
}

func checkRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, crmOn)
}

func setTimeout(w http.ResponseWriter, r *http.Request) {
	timeoutStr := r.URL.Query().Get("timeout")
	if timeoutStr != "" {
		newTimeout, err := strconv.Atoi(timeoutStr)
		if err != nil {
			http.Error(w, "Valor inválido", http.StatusBadRequest)
			return
		}

		crmTimeout = time.Duration(newTimeout) * time.Second

		stopRoutine(nil, nil)
		startRoutine(nil, nil)
	}
	fmt.Fprintln(w)
}
