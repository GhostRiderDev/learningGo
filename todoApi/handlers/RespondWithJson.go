package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, r *http.Request, code int, data any) {
	log.Printf("::: %v ::: %v %v", code, r.Method, r.URL)

	w.Header().Set("Content-Type", "application/json")

	d, err := json.Marshal(data)

	if err != nil {
		log.Println("---error marshaling data---")
		log.Printf("error: %v", err)
		log.Printf("data: %v", data)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write(nil); err != nil {
			log.Println(err)
		}
		return
	}

	w.WriteHeader(code)

	if _, err := w.Write(d); err != nil {
		log.Println(err)
	}
}