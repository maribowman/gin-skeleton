package model

import (
	"time"
)

type Trade struct {
	Date      time.Time
	Direction string
	ATR       float32
	Entry     float32
	BestDays  []float32
	WorstDays []float32
	MaxProfit []float32 // long: best - entry; short: entry - best -> set negative value to 0
	MaxLoss   []float32 // long: worst - entry; short: entry - worst -> set positive value to 0
}

type Trades []Trade

func (t Trades) Less(i, j int) bool { return t[i].Date.Before(t[j].Date) }
func (t Trades) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trades) Len() int           { return len(t) }

type Result struct {
	Symbol, ATR, SMA1, SMA2 string
	Trades                  Trades
	ProfitSum               []float32 // aggregates all profits per day per trade
	LossSum                 []float32 // aggregates all losses per day per trade
	Edge                    []float32 // sum(ProfitSum[x])/abs(sum(LossSum[x]))
}

// Well, CSV stands for comma separated values, if you open up your generated file through note pad you will have a real look what those data really looks like and you will have an idea on what your asking for.
//
//    to insert empty row you just do "","","",""
//    To insert empty column (let's say, first column is empty) you do "","data1","data2"
//    to insert your new table 2, you do the same as creating your table1 but you insert your table heads first after the table1. so the data should like this:
//
//column1-1,column1-2,column1-3
//datat1-1,"data1-2,data11-3
//column2-1,column2-2,column2-3
//data12-1,data12-2,data12-3"
