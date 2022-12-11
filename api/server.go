package api

import (
	"TradingBot/trader"
	"TradingBot/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	listenAddress string
	trader        *trader.Trader
}

func NewServer(listenAddress string) *Server {
	return &Server{
		trader:        trader.NewTrader(),
		listenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting Server...")

	http.HandleFunc("/healthCheck", s.handleHealthCheck)

	http.HandleFunc("/bot/positions", s.handleGetOpenPositions)
	http.HandleFunc("/bot/start", s.handleStartBot)
	http.HandleFunc("/bot/stop", s.handleStopBot)

	return http.ListenAndServe(s.listenAddress, nil)
}

func (s *Server) handleStartBot(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)

	if request.Method != http.MethodPut {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Starting bot...")

	s.trader.Start()
}

func (s *Server) handleStopBot(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)

	if request.Method != http.MethodPut {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Stopping bot...")

	s.trader.Stop()
}

func (s *Server) handleGetOpenPositions(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//TODO: This would return open positions
}

func (s *Server) handleHealthCheck(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	health := types.NewHealthCheck()

	err := json.NewEncoder(writer).Encode(health)

	if err != nil {
		panic("Error handling health check!")
	}
}

func (s *Server) requestLogging(request *http.Request) {
	println("Inbound Request: " + request.Method + " '" + request.RequestURI + "'")
}
