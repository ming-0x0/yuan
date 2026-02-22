package main

import (
	"fmt"
	"time"

	"github.com/godruoyi/go-snowflake"
	"github.com/ming-0x0/yuan/internal/modules/user/domain/user"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneICT()
	snowflake.SetStartTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local))

	user, _ := user.New("abc", "abc", "abc", "abc")

	fmt.Println("user: ", user)
}
