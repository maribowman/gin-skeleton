package model

import (
	"container/list"
	"math"
)

type MA struct {
	quote  float32
	size   int
	values *list.List
	spare  Candle
}

func NewMA(size int) *MA {
	return &MA{size: size, values: list.New()}
}

func (ma *MA) Tail(candle Candle) {
	ma.values.PushBack(candle)
	if ma.values.Len() > ma.size {
		ma.spare = ma.values.Front().Value.(Candle)
		ma.values.Remove(ma.values.Front())
	}
}

// TODO impl EMA and SMA separated on same value list

func (ma *MA) GetSMA() float32 {
	if ma.values.Len() != ma.size {
		return -1
	}
	var aggregation float32 = 0
	element := ma.values.Front()
	for index := 0; element != nil; index++ {
		aggregation += element.Value.(Candle).Close
		element = element.Next()
	}
	return aggregation / float32(ma.size)
}

// GetATR  is based on sma
func (ma *MA) GetATR() float32 {
	if ma.values.Len() != ma.size {
		return -1
	}
	var aggregation float32 = 0
	previous := ma.spare
	element := ma.values.Front()
	for index := 0; element != nil; index++ {
		aggregation += calcTR(element.Value.(Candle), previous)
		previous = element.Value.(Candle)
		element = element.Next()
	}
	// rounding 3 digits
	return float32(math.Round(float64(aggregation/float32(ma.size))*1000) / 1000)
}

func calcTR(current, previous Candle) float32 {
	// current high - low (the traditional range)
	tr := float32(math.Round(math.Abs(float64(current.High-current.Low))*100) / 100)
	if previous.Close == 0 {
		return tr
	}
	// current high - previous close
	higher := float32(math.Round(math.Abs(float64(current.High-previous.Close))*100) / 100)
	// previous close - current low
	lower := float32(math.Round(math.Abs(float64(previous.Close-current.Low))*100) / 100)
	if tr < higher {
		tr = higher
	}
	if tr < lower {
		tr = lower
	}
	return tr
}
