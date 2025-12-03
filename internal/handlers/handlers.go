package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/garylouisstewart/go-micro/internal/middleware"
)

type Response struct {
        Message   string  `json:"message"`
				Timestamp time.Time `json:"timestamp"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

   
	middleware.Logger.Debug("preparing hello response", "client_ip", r.RemoteAddr)

	resp := Response{
					Message:    "Hello World!",
					Timestamp: time.Now(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}


func HealthHandler(w http.ResponseWriter, r *http.Request) {
	      w.WriteHeader(http.StatusOK)
	      w.Write([]byte("OK"))
}

