package model

import (
	"time"
)

type Stock struct {
	Name              string    `gorm:"notNull"`
	Market            string    `gorm:"notNull"`
	YahooSymbol       string    `gorm:"notNull;unique"`
	TradingviewSymbol string    `gorm:"primaryKey"`
	ISIN              string    `gorm:"notNull;unique"`
	Start             time.Time `gorm:"notNull"`
}

type StockTrendOverview struct {
	Symbol  string `gorm:"primaryKey"`
	Daily   string
	Weekly  string
	Monthly string
}

type Sentiment struct {
	Time               time.Time `gorm:"primaryKey"`
	StockFearAndGreed  int       `gorm:"notNull"`
	CryptoFearAndGreed int       `gorm:"notNull"`
}

type AssetEntity struct {
	Time   time.Time `gorm:"primaryKey" json:"time"`
	Symbol string    `gorm:"notNull" json:"symbol"`
	Candle Candle    `gorm:"notNull" json:"candle"`
	//Indicators Indicators `gorm:"notNull" json:"indicators"` // todo TBD
}

type Indicators struct {
	ATR  float32 `json:"atr"`
	RSI  float32 `json:"rsi"`
	EMA  float32 `json:"ema"`
	MACD float32 `json:"macd"`
}
