package hw_1

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func Run() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(time.String())
}
