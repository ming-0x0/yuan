package app

import (
	"context"
	"time"

	"github.com/ming-0x0/yuan/internal/account/domain"
	"github.com/sony/sonyflake/v2"
	"github.com/sony/sonyflake/v2/awsutil"
)

var sf *sonyflake.Sonyflake

func init() {
	st := sonyflake.Settings{
		MachineID: awsutil.AmazonEC2MachineID,
		TimeUnit:  time.Millisecond,
		StartTime: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	var err error
	sf, err = sonyflake.New(st)
	if err != nil {
		panic(err)
	}
}

type AccountApp interface {
	CreateAccount(ctx context.Context, email string, password string) error
}

type accountApp struct {
	accountRepo domain.AccountRepository
}
