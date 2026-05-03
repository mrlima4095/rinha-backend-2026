package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/robogg133/rinha-backend-2026/internal/references"
	"github.com/robogg133/rinha-backend-2026/pkg/payments"
)

func main() {

	fmt.Println("Starting to init")
	if err := references.Init(); err != nil {
		panic(err)
	}
	fmt.Println("Successufully read references")

	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	http.HandleFunc("/fraud-score", func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf("[%s] Received request\n", time.Now().String())

		request := new(payments.Payment)

		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			return
		}

		vecptr := request.ToVector()

		fmt.Printf("[Request] %s\n", vecptr.Stringify())

		score := FindFraudScore(*vecptr)
		approved := "false"

		if score < 0.6 {
			approved = "true"
		}

		answer := fmt.Sprintf("{\"approved\": %s, \"fraud_score\": %.1f}", approved, score)
		fmt.Printf("[Response] %s\n", answer)
		fmt.Fprint(w, answer)
	})

	http.ListenAndServe(":8080", nil)
}
