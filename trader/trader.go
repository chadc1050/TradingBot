package trader

import (
	"TradingBot/auth"
	"TradingBot/config"
	"context"
	"fmt"
	"github.com/zricethezav/go-tdameritrade"
)

type Trader struct {
	Authentication *auth.Authentication
	BotConfig      *config.BotConfig
}

func NewTrader(config config.BotConfig) *Trader {

	fmt.Println("Creating Trader...")

	return &Trader{
		Authentication: auth.NewAuthentication(),
	}
}

func (t *Trader) GetAccountDetails(accountId string, positions bool, orders bool) (*tdameritrade.Account, error) {
	client, err := t.Authentication.GetAuthenticatedClient()
	if err != nil {
		return nil, err
	}

	account, _, err := client.Account.GetAccount(context.Background(), accountId, &tdameritrade.AccountOptions{Position: positions, Orders: orders})

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (t *Trader) Start() {
	fmt.Println("Starting Bot...")
}

func (t *Trader) Stop() {
	fmt.Println("Stopping Bot...")
}

func (t *Trader) process() {
	// Function where all trading processes will be endlessly looped through
	// Check if in trading hours
	// Close positions
	// Check if there is unallocated buying power
	// Position discovery
}
