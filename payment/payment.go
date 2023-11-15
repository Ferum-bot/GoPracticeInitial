package main

import (
	"log"
	"sync"
	"time"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once
}

func (p *Payment) Process() {
	timestamp := time.Now()
	process := func() {
		p.process(timestamp)
	}
	p.once.Do(process)
}

func (p *Payment) process(timestamp time.Time) {
	formattedTimestamp := timestamp.Format(time.RFC3339)
	log.Printf("[%s] %s -> %.2f -> %s", formattedTimestamp, p.From, p.Amount, p.To)
}

func main() {
	payment := Payment{
		From:   "Matvey",
		To:     "Stas",
		Amount: 12.32,
	}

	payment.Process()
	payment.Process()
}
