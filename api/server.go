package api

import (
	"TradingBot/config"
	"TradingBot/trader"
	"TradingBot/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	trader        *trader.Trader
	listenAddress string
}

func NewServer(listenAddress string) *Server {

	botConfig := config.BotConfig{
		MaxPositionAllocationPercent: 5,
	}

	return &Server{
		trader:        trader.NewTrader(botConfig),
		listenAddress: listenAddress,
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting Server...")

	http.HandleFunc("/health-check", s.handleHealthCheck)

	// Called by TDA
	http.HandleFunc("/tda/auth", s.handleTDAAuthCallback)

	http.HandleFunc("/bot/auth", s.handleUserAuth)
	http.HandleFunc("/bot/is-logged-in", s.handleIsLoggedIn)
	http.HandleFunc("/bot/positions", s.handleGetOpenPositions)
	http.HandleFunc("/bot/start", s.handleStartBot)
	http.HandleFunc("/bot/stop", s.handleStopBot)
	http.HandleFunc("/bot/account-details/account/", s.handleGetAccountDetails)

	if err := http.ListenAndServe(s.listenAddress, nil); err != nil {
		fmt.Println("Error occurred starting server!")
		return err
	}

	println("Server Ready")
	return nil
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

	if err := json.NewEncoder(writer).Encode(health); err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error handling health check!")
		return
	}
}

func (s *Server) handleUserAuth(writer http.ResponseWriter, request *http.Request) {
	s.requestLogging(request)
	s.configureCors(&writer)

	if request.Method != http.MethodPost {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	redirectUrl, err := s.trader.Authentication.Authenticate(writer, request)

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	authUrl := types.AuthUrlResponse{AuthUrl: redirectUrl}

	if err := json.NewEncoder(writer).Encode(authUrl); err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleTDAAuthCallback(writer http.ResponseWriter, request *http.Request) {
	s.requestLogging(request)
	s.configureCors(&writer)

	// TODO: Get hostname from auth url
	if request.Header.Get("Origin") != os.Getenv("TDA_AUTH_URL") {
		fmt.Println("WARN: Origin of request was not TDA Auth!")
	}

	if err := s.trader.Authentication.Callback(writer, request); err != nil {
		fmt.Println("Failed to authenticate account!")
		return
	}

	fmt.Println("Successfully authenticated account!")
}

func (s *Server) handleIsLoggedIn(writer http.ResponseWriter, request *http.Request) {

	s.requestLogging(request)
	s.configureCors(&writer)

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	status, err := s.trader.Authentication.IsLoggedIn(request)

	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	loggedIn := types.LoggedIn{
		IsLoggedIn: status,
	}

	if err := json.NewEncoder(writer).Encode(loggedIn); err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error occurred writing response for logged in!")
		return
	}
}

func (s *Server) handleGetAccountDetails(writer http.ResponseWriter, request *http.Request) {
	s.requestLogging(request)
	s.configureCors(&writer)

	if request.Method != http.MethodGet {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	account := strings.TrimPrefix(request.URL.Path, "/bot/account-details/account/")

	accountResponse, err := s.trader.GetAccountDetails(account, true, true)

	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Error occurred retrieving account details", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(writer).Encode(accountResponse); err != nil {
		fmt.Println("Error occurred serving account details!")
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) configureCors(writer *http.ResponseWriter) {
	(*writer).Header().Set("Access-Control-Allow-Origin", "*")
}

func (s *Server) requestLogging(request *http.Request) {
	println("Inbound Request: " + request.Method + " '" + request.RequestURI + "'")
}
