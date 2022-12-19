package main

import (
	"TradingBot/api"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Print("  _____              _ _               ____        _   \n" +
		" |_   _| __ __ _  __| (_)_ __   __ _  | __ )  ___ | |_ \n" +
		"   | || '__/ _` |/ _` | | '_ \\ / _` | |  _ \\ / _ \\| __|\n" +
		"   | || | | (_| | (_| | | | | | (_| | | |_) | (_) | |_ \n" +
		"   |_||_|  \\__,_|\\__,_|_|_| |_|\\__, | |____/ \\___/ \\__|\n" +
		"                               |___/                   \n")

	fmt.Println("Starting Trading Bot Server...")

	fmt.Println("--------------------------------")
	fmt.Println("TD Ameritrade Consumer Key: " + os.Getenv("TDA_CONSUMER_KEY"))
	fmt.Println("TD Ameritrade Auth URL: " + os.Getenv("TDA_AUTH_URL"))
	fmt.Println("TD Ameritrade Redirect URL: " + os.Getenv("TDA_REDIRECT_URL"))
	fmt.Println("--------------------------------")

	listenAddress := flag.String("listenAddress", ":8080", "Server Address")

	flag.Parse()

	server := api.NewServer(*listenAddress)

	err := server.Start()
	if err != nil {
		fmt.Println("Error starting server!")
	}
}
