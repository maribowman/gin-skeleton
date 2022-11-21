package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Candle struct {
	Date    time.Time `json:"date"`
	Current float32   `json:"-"`
	Open    float32   `json:"open"`
	High    float32   `json:"high"`
	Low     float32   `json:"low"`
	Close   float32   `json:"close"`
	Volume  int       `json:"volume"`
}

// https://medium.com/iostrap/manage-postgresql-json-data-with-go-golang-b0ae416972c5
func (Candle) GormDataType() string {
	return "JSONB"
}

func (Candle) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "JSONB"
}

// Scan scans value into Jsonb, implements sql.Scanner interface
func (candle *Candle) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	var result Candle
	err := json.Unmarshal(bytes, &result)
	*candle = result
	return err
}

// Value returns json value, implements driver.Valuer interface
func (candle Candle) Value() (driver.Value, error) {
	if candle == (Candle{}) {
		return nil, nil
	}
	return json.Marshal(candle)
}
