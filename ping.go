package main

import 	(
	"github.com/go-ping/ping"
	"fmt"
)

func p() {
	pinger, err := ping.NewPinger("127.0.0.1")

	if err != nil { panic(err)}

	pinger.Count = 3
	err = pinger.Run()

	if err != nil { panic(err)}

	fmt.Println(pinger.Statistics())
}
