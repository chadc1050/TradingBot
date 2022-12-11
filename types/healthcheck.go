package types

import (
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"strconv"
)

type HealthCheck struct {
	Status               string `json:"status"`
	OS                   string `json:"OS"`
	Architecture         string `json:"architecture"`
	AvailableMemory      string `json:"availableMemory"`
	MemoryUsed           string `json:"memoryUsed"`
	MaxMemory            string `json:"maxMemory"`
	PercentageMemoryUsed string `json:"percentageMemoryUsed"`
}

func NewHealthCheck() *HealthCheck {

	memory, err := mem.VirtualMemory()

	if err != nil {
		panic("Error occurred retrieving memory!")
	}

	return &HealthCheck{
		Status:               "RUNNING",
		OS:                   runtime.GOOS,
		Architecture:         runtime.GOARCH,
		AvailableMemory:      strconv.FormatUint(memory.Free/(1024*1024), 10) + "MB",
		MaxMemory:            strconv.FormatUint(memory.Total/(1024*1024), 10) + "MB",
		MemoryUsed:           strconv.FormatUint(memory.Used/(1024*1024), 10) + "MB",
		PercentageMemoryUsed: strconv.FormatFloat(memory.UsedPercent, 'f', 0, 32) + "%",
	}
}
