package main

import (
	"fmt"

	"github.com/ming-0x0/yuan/internal/common/domain/id"
	"github.com/ming-0x0/yuan/pkg/timezone"
)

func main() {
	timezone.SetTimeZoneICT()
	fmt.Println(id.MustNew())
}
