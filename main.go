package main

import (
	"fmt"
	"time"

	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneUTC()
	fmt.Println(time.Now())
}
