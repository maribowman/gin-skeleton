package model

import "time"

type EdgeDTO struct {
	Symbol     string `json:"symbol"`
	Entry      string `json:"entry"`
	EdgeTrades []struct {
		Date      string `json:"date"`
		Direction string `json:"direction"`
	} `json:"trades"`
}

type EdgeTrade struct {
	Date      time.Time
	Direction string
}
