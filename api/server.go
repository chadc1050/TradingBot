package api

import (
	"TradingBot/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	listenAddress string
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting Server...")

	http.HandleFunc("/healthCheck", s.handleHealthCheck)

	return http.ListenAndServe(s.listenAddress, nil)
}

func (s *Server) handleHealthCheck(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	health := types.NewHealthCheck("RUNNING")

	err := json.NewEncoder(writer).Encode(health)

	if err != nil {
		panic("Error handling health check!")
	}

}

func (s *Server) requestLogging(request *http.Request) {
	println("Inbound Request: " + request.Method + " '" + request.RequestURI + "'")
}
