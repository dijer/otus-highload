package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type UserDate time.Time

const layout = "2006-01-02"

func (d *UserDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*d).Format(layout))
}

func (d *UserDate) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"`+layout+`"`, string(data))
	if err != nil {
		return err
	}

	*d = UserDate(parsedTime)
	return nil
}

func (d *UserDate) Scan(value interface{}) error {
	if value == nil {
		*d = UserDate(time.Time{})
		return nil
	}

	switch v := value.(type) {
	case string:
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			return err
		}
		*d = UserDate(parsedTime)
	case time.Time:
		*d = UserDate(v)
	default:
		return fmt.Errorf("unsupported type for UserDate %T", v)
	}

	return nil
}

func (d UserDate) Value() (driver.Value, error) {
	return time.Time(d).Format(layout), nil
}
