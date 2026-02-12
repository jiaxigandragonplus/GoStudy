package test

import (
	"log"
	"testing"
	"time"
)

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func TestTimer(t *testing.T) {
	TickerDemo()
}
