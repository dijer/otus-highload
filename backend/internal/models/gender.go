package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g *Gender) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	parsed := Gender(s)

	switch parsed {
	case Male, Female:
		*g = parsed
		return nil
	default:
		return fmt.Errorf("invalid gender %s", s)
	}
}

func (g Gender) Value() (driver.Value, error) {
	if g == "" {
		return nil, nil
	}

	return string(g), nil
}
