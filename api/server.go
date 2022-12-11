package api

import (
	"TradingBot/trader"
	"TradingBot/types"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	trader        *trader.Trader
	listenAddress string
}

func NewServer(listenAddress string) *Server {
	return &Server{
		trader:        trader.NewTrader(),
		listenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting Server...")

	http.HandleFunc("/health-check", s.handleHealthCheck)

	http.HandleFunc("/tda/auth", s.handleTDAAuthCallback)
	http.HandleFunc("/tda/is-logged-in", s.handleIsLoggedIn)
	http.HandleFunc("/tda/log-in", s.handleLogIn)

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

func (s *Server) handleTDAAuthCallback(writer http.ResponseWriter, request *http.Request) {
	s.requestLogging(request)

	if request.Header.Get("Origin") != "https://auth.tdameritrade.com" {
		fmt.Println("WARN: Origin of request was not TDA Auth!")
	}

	refreshToken := request.URL.Query().Get("code")

	if len(refreshToken) == 0 {
		fmt.Println("Refresh token not present in request!")
	}

}

func (s *Server) handleIsLoggedIn(writer http.ResponseWriter, request *http.Request) {
	// Will check if code in client is set to a value
}

func (s *Server) handleLogIn(writer http.ResponseWriter, request *http.Request) {
	// Handles TDA OAuth and will set the code
}

func (s *Server) requestLogging(request *http.Request) {
	println("Inbound Request: " + request.Method + " '" + request.RequestURI + "'")
}
