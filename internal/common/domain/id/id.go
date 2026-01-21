package id

import (
	"time"

	"github.com/sony/sonyflake/v2"
	"github.com/sony/sonyflake/v2/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	var err error
	sf, err = sonyflake.New(sonyflake.Settings{
		StartTime: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		MachineID: awsutil.AmazonEC2MachineID,
	})
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
	id, err := New()
	if err != nil {
		panic(err)
	}
	return id
}
