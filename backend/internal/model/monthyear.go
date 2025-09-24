package model

import (
	"database/sql/driver"
	"fmt"

	"time"
)

const monthYearLayout = "01-2006"

type MonthYear struct {
	time.Time
}

func (m *MonthYear) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	t, err := time.Parse(monthYearLayout, s)
	if err != nil {
		return fmt.Errorf("invalid month-year format: %w", err)
	}
	m.Time = t
	return nil
}

func (m *MonthYear) UnmarshalText(text []byte) error {
	s := string(text)
	t, err := time.Parse(monthYearLayout, s)
	if err != nil {
		return fmt.Errorf("invalid month-year format: %w", err)
	}
	m.Time = t
	return nil
}


func (m MonthYear) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", m.Time.Format(monthYearLayout))), nil
}

func (m MonthYear) Value() (driver.Value, error) {
	return m.Time, nil
}

func (m *MonthYear) Scan(value interface{}) error {
	if value == nil {
		m.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		m.Time = v
		return nil
	default:
		return fmt.Errorf("cannot scan %T into MonthYear", value)
	}
}
