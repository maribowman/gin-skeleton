package model

import (
	"container/list"
	"fmt"
	"time"
)

const (
	Hourly = "H1"
	Daily  = "D1"
	Weekly = "W1"
)

type timeKeys map[string]*list.Element

type OrderedCandleHistory struct {
	timeFrame  string
	keys       timeKeys
	linkedList *list.List
}

func NewOrderedCandleHistory(timeFrame string, candles ...Candle) *OrderedCandleHistory {
	history := OrderedCandleHistory{
		timeFrame,
		make(timeKeys),
		list.New(),
	}
	for _, candle := range candles {
		history.Add(candle)
	}
	return &history
}

func (history *OrderedCandleHistory) createKey(time time.Time) string {
	year, month, day := time.Date()
	hour, _, _ := time.Clock()
	switch history.timeFrame {
	case Hourly:
		return fmt.Sprintf("%d-%d-%d-%d", year, int(month), day, hour)
	case Daily:
		return fmt.Sprintf("%d-%d-%d", year, int(month), day)
	default:
		return ""
	}
}

func (history *OrderedCandleHistory) Add(candle Candle) {
	key := history.createKey(candle.Date)
	_, exists := history.keys[key]
	if !exists {
		history.keys[key] = history.linkedList.PushBack(candle)
	}
}

func (history *OrderedCandleHistory) Remove(date time.Time) {
	key := history.createKey(date)
	element, exists := history.keys[key]
	if exists {
		history.linkedList.Remove(element)
		delete(history.keys, key)
	}
}

func (history *OrderedCandleHistory) First() Candle {
	return history.linkedList.Front().Value.(Candle)
}

func (history *OrderedCandleHistory) Has(date time.Time) bool {
	_, exists := history.keys[history.createKey(date)]
	return exists
}

func (history *OrderedCandleHistory) HasNext(candle Candle) bool {
	return !history.Next(candle.Date).Date.IsZero()
}

func (history *OrderedCandleHistory) Next(date time.Time) Candle {
	element, exists := history.keys[history.createKey(date)]
	if !exists || element.Next() == nil {
		return Candle{}
	}
	return element.Next().Value.(Candle)
}

func (history *OrderedCandleHistory) HasPrevious(candle Candle) bool {
	return !history.Previous(candle.Date).Date.IsZero()
}

func (history *OrderedCandleHistory) Previous(date time.Time) Candle {
	element, exists := history.keys[history.createKey(date)]
	if !exists || element.Prev() == nil {
		return Candle{}
	}
	return element.Prev().Value.(Candle)
}

func (history *OrderedCandleHistory) Get(date time.Time) Candle {
	element, exists := history.keys[history.createKey(date)]
	if !exists {
		return Candle{}
	}
	return element.Value.(Candle)
}

func (history *OrderedCandleHistory) GetRange(from, to time.Time) []Candle {
	candleRange := []Candle{}
	candle := history.Get(from)
	if candle.Date.IsZero() {
		return candleRange
	}
	candleRange = append(candleRange, candle)
	for history.HasNext(candle) {
		candle = history.Next(candle.Date)
		candleRange = append(candleRange, candle)
		if candle.Date == to {
			// TODO check if half full range is ok
			break
		}
	}
	return candleRange
}

func (history *OrderedCandleHistory) GetPrevious(date time.Time, amount int) []Candle {
	result := []Candle{}
	element, exists := history.keys[history.createKey(date)]
	if !exists {
		return result
	}
	for i := 0; i < amount; i++ {
		element = element.Prev()
		if element == nil {
			break
		}
		result = append(result, element.Value.(Candle))
	}
	return result
}

func (history *OrderedCandleHistory) IsLast(candle Candle) bool {
	_, exists := history.keys[history.createKey(candle.Date)]
	if exists {
		return candle == history.linkedList.Back().Value.(Candle)
	}
	return false
}

func (history *OrderedCandleHistory) Size() int {
	return len(history.keys)
}

func (history *OrderedCandleHistory) Values() []Candle {
	values := make([]Candle, len(history.keys))
	element := history.linkedList.Front()
	for index := 0; element != nil; index++ {
		values[index] = element.Value.(Candle)
		element = element.Next()
	}
	return values
}
