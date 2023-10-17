package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dgmt/parcel_master/internal/parcelcalc"
	"github.com/go-chi/chi/v5"
)

func main() {

	router := chi.NewRouter()

	router.Route("/api/v1/parcelcalc", func(r chi.Router) {
		r.Get("/{orderAmt}", ParcelController)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println(err)
	}
}

func ParcelController(w http.ResponseWriter, r *http.Request) {
	orderAmtString := chi.URLParam(r, "orderAmt")
	orderAmt, err := strconv.ParseInt(orderAmtString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := parcelcalc.ParcelCalculator(orderAmt)
	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
