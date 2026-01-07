package sonyflake

import (
	"time"

	"github.com/sony/sonyflake/v2"
)

type Sonyflake struct {
	*sonyflake.Sonyflake
}

func New() (*Sonyflake, error) {
	var st sonyflake.Settings

	st.MachineID = func() (int, error) {
		return 1, nil
	}
	st.TimeUnit = time.Millisecond
	st.StartTime = time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	sf, err := sonyflake.New(st)
	if sf == nil {
		return nil, err
	}

	return &Sonyflake{Sonyflake: sf}, nil
}

func (s *Sonyflake) Next() (int64, error) {
	return s.NextID()
}
