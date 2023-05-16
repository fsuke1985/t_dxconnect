package main

import (
	"fmt"
	"os"
	"github.com/go-ping/ping"
)

func p() {
	pinger, err := ping.NewPinger(os.Getenv("PINGTARGET"))

	if err != nil { panic(err)}

	pinger.Count = 3
	err = pinger.Run()

	if err != nil { panic(err)}

	fmt.Println(pinger.Statistics())
}
