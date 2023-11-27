package route

import (
	"fmt"
	"log"
	"manuall/routine-api/controller"
	"net/http"
	"time"
)

var (
	ticker *time.Ticker
	quit   chan struct{}
)

func Crm() {
	quit = make(chan struct{})

	http.HandleFunc("/crm/on", startRoutine)
	http.HandleFunc("/crm/off", stopRoutine)
	http.HandleFunc("/crm/check", checkRoutine)

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func startRoutine(w http.ResponseWriter, r *http.Request) {
	if ticker == nil {
		ticker = time.NewTicker(15 * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					controller.Crm()
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	}
	fmt.Fprintln(w)
}

func stopRoutine(w http.ResponseWriter, r *http.Request) {
	if ticker != nil {
		close(quit)
		ticker = nil
		quit = make(chan struct{})
	}
	fmt.Fprintln(w)
}

func checkRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w)
}
