package main

import (
	"fmt"
	"time"

	"github.com/godruoyi/go-snowflake"
	"github.com/ming-0x0/yuan/internal/modules/iam/domain/user"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneICT()
	snowflake.SetStartTime(time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local))

	user, _ := user.New("abc", "abc@gmail.com", "abc", "abc", true)

	fmt.Println("user: ", user)
}
