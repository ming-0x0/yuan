package app

import (
	"time"

	"github.com/ming-0x0/yuan/internal/account/domain/account"
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

type AccountApp struct {
	accountRepo account.AccountRepository
}

func New(
	accountRepo account.AccountRepository,
) *AccountApp {
	return &AccountApp{
		accountRepo: accountRepo,
	}
}
