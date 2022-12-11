package main

import (
	"TradingBot/api"
	"flag"
	"fmt"
)

func main() {

	fmt.Print("  _____              _ _               ____        _   \n" +
		" |_   _| __ __ _  __| (_)_ __   __ _  | __ )  ___ | |_ \n" +
		"   | || '__/ _` |/ _` | | '_ \\ / _` | |  _ \\ / _ \\| __|\n" +
		"   | || | | (_| | (_| | | | | | (_| | | |_) | (_) | |_ \n" +
		"   |_||_|  \\__,_|\\__,_|_|_| |_|\\__, | |____/ \\___/ \\__|\n" +
		"                               |___/                   \n")
	fmt.Println("Starting Trading Bot Server...")

	listenAddress := flag.String("listenAddress", ":443", "Server Address")

	flag.Parse()

	server := api.NewServer(*listenAddress)

	err := server.Start()
	if err != nil {
		fmt.Println("Error starting server!")
	}
}
