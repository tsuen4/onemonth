package onemonth

import (
	"fmt"
	"time"
)

// Month has BeginDay and EndDay.
type Month struct {
	BeginDay time.Time
	EndDay   time.Time
}

// New returns a Month struct with BeginDay and EndDay.
func New(year, month int) (*Month, error) {
	if month < 1 || 12 < month {
		return nil, fmt.Errorf("month must be between 1 and 12, got: %d", month)
	}

	begin := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	nextMonth := int(begin.Month()) + 1
	end := time.Date(year, time.Month(nextMonth), 0, 0, 0, 0, 0, time.UTC)

	return &Month{
		BeginDay: begin,
		EndDay:   end,
	}, nil
}

// Iterate through the days of the Month.
func (m Month) Iterate(f func(day time.Time)) {
	for day := m.BeginDay; !day.After(m.EndDay); day = day.AddDate(0, 0, 1) {
		f(day)
	}
}

// Month returns the time.Month.
func (m Month) Month() time.Month {
	return m.BeginDay.Month()
}
