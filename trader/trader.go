package trader

import (
	"github.com/chadc1050/TDAClient/client"
	"os"
)

type Runnable interface {
	Start()
	Stop()
}

type Trader struct {
	client *client.TDAClient
	active bool
}

func NewTrader() *Trader {
	return &Trader{
		client: client.NewClient(os.Getenv("ACCOUNT_ID"), os.Getenv("ACCOUNT_KEY")),
		active: false,
	}
}

func (t *Trader) Start() {
	if !t.active {

	}
}

func (t *Trader) Stop() {
	if t.active {

	}
}

func (t *Trader) process() {
	// Function where all trading processes will be endlessly looped through
	// Check if in trading hours
	// Close positions
	// Check if there is unallocated buying power
	// Position discovery
}
