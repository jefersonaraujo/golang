package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCepHander)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHander(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

//http://localhost:8080/?cep=60766320
