package id

import (
	"time"

	"github.com/sony/sonyflake/v2"
)

var sf *sonyflake.Sonyflake

func init() {
	st := sonyflake.Settings{
		StartTime: time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local),
	}
	var err error
	sf, err = sonyflake.New(st)
	if err != nil {
		panic(err)
	}
}

type ID int64

func New() (ID, error) {
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return ID(id), nil
}

func MustNew() ID {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return ID(id)
}
