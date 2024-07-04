package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

type VerifyQuestion struct {
	Q1 *string `json:"q1"`
	A1 *string `json:"a1"`
	Q2 *string `json:"q2"`
	A2 *string `json:"a2"`
	Q3 *string `json:"q3"`
	A3 *string `json:"a3"`
}

func (c *VerifyQuestion) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}
func (c VerifyQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
