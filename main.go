package main

import (
	"TradingBot/api"
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting Trading Bot...")

	listenAddress := flag.String("listenAddress", ":8080", "Server Address")

	flag.Parse()

	server := api.NewServer(*listenAddress)

	err := server.Start()
	if err != nil {
		fmt.Println("Error starting server!")
	}
}
